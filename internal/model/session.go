package model

type Session struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	SessionID string `json:"session_id"`
	UserID    uint   `json:"user_id"`
	User      User   `gorm:"foreignKey:UserID" json:"user"`
}
type SessionData struct {
	Session Session `json:"session"`
	Message string  `json:"message"`
}
