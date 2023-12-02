package entity

type User struct {
	Id       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
}
