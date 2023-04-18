package structures

type UserRole struct {
	UserId int `form:"userId" json:"userId,omitempty"`
	RoleId int `form:"roleId" json:"roleId,omitempty"`
}
type UserSocNetwork struct {
	UserId        int    `form:"userId" json:"userId,omitempty"`
	NetworkId     string `form:"networkId" json:"roleId,omitempty"`
	NetworkNameId int    `form:"networkId" json:"networkNameId"`
}
