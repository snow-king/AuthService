package models

type Roles struct {
	ID        int    `gorm:"index:,unique" json:"ID,omitempty"`
	Name      string `gorm:"name" json:"name,omitempty"`
	ShortName string `gorm:"shortName" json:"shortName,omitempty"`
	JwtName   string `gorm:"JwtName" json:"jwtName"`
}

func (Roles) TableName() string {
	return "roles"
}
