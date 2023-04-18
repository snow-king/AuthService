package SocialService

type SocService interface {
	Index() string
	Callback(code string) (string, error)
}
type Context struct {
	service SocService
}

func (c *Context) Service(a SocService) {
	c.service = a
}
