package dto

//ProductUpdateDTO is used when client PUT or update data product
type ProductUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Type        string `json:"type" form:"type" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	PriceBefore int    `json:"price_before" form:"price_before" binding:"required"`
	PriceToday  int    `json:"price_today" form:"price_today" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

//ProductCreateDTO is used when client POST or creating data product
type ProductCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Type        string `json:"type" form:"type" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	PriceBefore int    `json:"price_before" form:"price_before" binding:"required"`
	PriceToday  int    `json:"price_today" form:"price_today" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
