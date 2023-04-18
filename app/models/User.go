package models

type LoginUser struct {
	Login    string `form:"login" json:"login,omitempty" binding:"required"`
	Password string `form:"password" json:"password,omitempty" binding:"required"`
	Role     string `form:"role" json:"role" binding:"required"`
}
type User struct {
	Id           int            `gorm:"id"`
	NdsLogin     string         `gorm:"nds_login" json:"nds_login,omitempty" validator:"required,email"`
	Password     string         `json:"password,omitempty" validator:"required"`
	Roles        []Roles        `gorm:"many2many:user_roles; foreignKey:id; joinForeignKey:user_id;References:id; joinReferences:rol_id"`
	UserNetworks []UserNetworks `gorm:"foreignKey:user_id; references:id"`
}

func (User) TableName() string {
	return "users"
}
