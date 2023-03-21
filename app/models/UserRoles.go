package models

type UserRoles struct {
	ID     int `json:"ID,omitempty"`
	UserId int `gorm:"user_id" json:"userId,omitempty"`
	RolId  int `gorm:"rol_id" json:"rolId,omitempty"`
}

func (UserRoles) TableName() string {
	return "user_roles"
}
