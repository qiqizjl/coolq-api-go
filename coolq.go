package coolq

//CoolQ Base
type CoolQ struct {
	URL   string
	Token string
}

//Map is a map
type Map map[string]interface{}

//NewCoolQ 获得一个CoolQ
func NewCoolQ(url, token string) *CoolQ {
	coolq := &CoolQ{
		URL:   url,
		Token: token,
	}
	return coolq
}
