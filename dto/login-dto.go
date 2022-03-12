package dto

//LoginDTO is used when client post data from /login url
type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:8"`
}