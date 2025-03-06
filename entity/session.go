package entity

import "github.com/google/uuid"

type Session struct {
	UserId    uuid.UUID
	Token     string
	IPAddress string
	UserAgent string
	DeviceId  string
}
