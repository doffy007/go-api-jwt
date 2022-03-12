package entity

//User struct is represent users table in database
type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"` //->;<-;not null can write or read
	Token    string `gorm:"-" json:"token,omitempty"` //Token not save on database
	Products *[]Product `json:"products,omitempty"`
}
