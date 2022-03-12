package dto

//RegisterDTO is used when client post data from /register url
type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min:1"`
	Email    string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Password string `json:"password" form:"password" validate:"min:8" binding:"required"`
}
