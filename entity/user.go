package entity

type User struct {
	ID       uint   `gorm: "primary_key" json:"id"`
	UserName string `gorm:"unique;not null" json:"username"`
}
