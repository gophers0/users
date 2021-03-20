package model

const (
	UserRoleAdmin = "admin"
	UserRoleUser  = "user"
)

type User struct {
	Model
	Login    string `json:"login" gorm:"index"`
	Password string `json:"-"`
	Role     string `json:"role"`
}
