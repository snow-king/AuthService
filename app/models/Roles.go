package models

type Roles struct {
	ID        int    `gorm:"index:,unique" json:"ID,omitempty"`
	Name      string `gorm:"name" json:"name,omitempty"`
	ShortName string `gorm:"id" json:"shortName,omitempty"`
}

func (Roles) TableName() string {
	return "roles"
}
