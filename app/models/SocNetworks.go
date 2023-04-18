package models

type UserNetworks struct {
	Id            int    `gorm:"id"`
	UserID        int    `gorm:"user_id"`
	NetworkNameId int    `gorm:"network_name_id"`
	NetworkId     string `gorm:"network_id"`
}

func (UserNetworks) TableName() string {
	return "user_networks"
}

type SocNetworks struct {
	Id   int
	Name string
}

func (SocNetworks) TableName() string {
	return "soc_networks"
}
