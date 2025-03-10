package model

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRes struct {
	JwtToken string `json:"jwtToken"`
}
