package models

type LoginUser struct {
	Login    string `form:"login" json:"login,omitempty"`
	Password string `form:"password" json:"password,omitempty"`
}
type User struct {
	Id       int    `gorm:"id"`
	NdsLogin string `gorm:"nds_login" json:"nds_login,omitempty"`
	Password string `json:"password,omitempty"`
}

func (User) TableName() string {
	return "users"
}
