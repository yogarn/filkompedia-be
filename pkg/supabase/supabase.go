package supabase

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	storage_go "github.com/supabase-community/storage-go"
)

type Supabase struct {
	client storage_go.Client
}

type ISupabase interface {
	UploadFile(file *multipart.FileHeader, dir string) (string, error)
}

func New() ISupabase {
	url := fmt.Sprintf("%s/storage/v1", os.Getenv("SUPABASE_PROJECT_URL"))
	client := storage_go.NewClient(url, os.Getenv("SUPABASE_KEY"), nil)

	return Supabase{
		client: *client,
	}
}

func (s Supabase) UploadFile(file *multipart.FileHeader, dir string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	buffer := make([]byte, 512)
	_, err = src.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)

	path := dir + "/" + uuid.NewString() + filepath.Ext(file.Filename)

	_, err = s.client.UploadFile(
		os.Getenv("SUPABASE_BUCKET_NAME"),
		path,
		src,
		storage_go.FileOptions{
			ContentType: &contentType,
		},
	)

	if err != nil {
		return "", err
	}

	publicURL := fmt.Sprintf("%s/storage/v1/object/public/%s/%s",
		os.Getenv("SUPABASE_PROJECT_URL"),
		os.Getenv("SUPABASE_BUCKET_NAME"),
		path,
	)

	return publicURL, nil
}
