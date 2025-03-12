package model

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
	JwtToken string `json:"jwtToken"`
}
