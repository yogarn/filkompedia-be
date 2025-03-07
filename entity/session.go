package entity

import "github.com/google/uuid"

type Session struct {
	UserId    uuid.UUID `db:"user_id"`
	Token     string    `db:"token"`
	IPAddress string    `db:"ip_address"`
	UserAgent string    `db:"user_agent"`
	DeviceId  string    `db:"device_id"`
}
