package dto

//UserUpdateDTO is used when client PUT or update data profile
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:8"`
}


