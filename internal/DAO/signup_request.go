package DAO

type SignUpRequest struct {
	Email	 string  `json:"email"`
	Password string  `json:"password"`
}