package model

import (
	"time"
)

type RegisterReq struct {
	Username string `json:"username" db:"username" validate:"required,lte=32"`
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"password" db:"password" validate:"required,gte=8"`
}

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}

type LoginRes struct {
	JwtToken     string `json:"jwtToken"`
	RefreshToken string `json:"refreshToken"`
}

type SessionsRes struct {
	IPAddress string    `json:"ip_address"`
	ExpiresAt time.Time `json:"expires_at"`
	UserAgent string    `json:"user_agent"`
	DeviceId  string    `json:"device_id"`
}
