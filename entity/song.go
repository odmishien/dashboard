package entity

type Song struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	UserID uint   `gorm:"not null" json:"user_id"`
	Name   string `gorm:"not null" json:"name"`
	Artist string `gorm:"not null" json:"artist"`
}
