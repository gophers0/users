package model

type Session struct {
	Model
	User   User   `json:"user"`
	UserID uint   `json:"-" gorm:"index:common_session_idx"`
	Token  string `json:"token" gorm:"index:common_session_idx"`
}
