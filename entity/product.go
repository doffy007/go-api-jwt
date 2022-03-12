package entity

//Product struct is an represent products table on database
type Product struct {
	ID          uint64 `gorm:"primaryKey:autoIncrement" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Type        string `gorm:"type:varchar(255)" json:"type"`
	Description string `gorm:"type:text" json:"description"`
	PriceBefore int    `gorm:"type:int" json:"price_before"`
	PriceToday  int    `gorm:"type:int" json:"price_today"`
	UserID      uint64 `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
}
