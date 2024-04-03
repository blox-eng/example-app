package model

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
}

type UserData struct {
	User    User   `json:"user"`
	Message string `json:"message"`
}
