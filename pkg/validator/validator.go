package validator

import (
	"mime/multipart"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

func RegisterValidator(v *validator.Validate) {
	v.RegisterValidation("rfc3339date", Date)
	v.RegisterValidation("image_type", ImageType)
	v.RegisterValidation("image_size", ImageSize)
}

func Date(fl validator.FieldLevel) bool {
	_, err := time.Parse(time.RFC3339, fl.Field().String())
	return err == nil
}

func ImageType(fl validator.FieldLevel) bool {
	fileHeader, ok := fl.Field().Interface().(multipart.FileHeader)
	if !ok {
		return false
	}

	file, err := fileHeader.Open()
	if err != nil {
		return false
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return false
	}

	mimeType := http.DetectContentType(buffer)
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
	}

	return allowedTypes[mimeType]
}

func ImageSize(fl validator.FieldLevel) bool {
	fileHeader, ok := fl.Field().Interface().(multipart.FileHeader)
	if !ok {
		return false
	}

	const maxSize = 2 * 1024 * 1024
	return fileHeader.Size <= maxSize
}
