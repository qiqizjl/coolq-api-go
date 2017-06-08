package coolq

//
type CoolQ struct {
	URL   string
	Token string
}

type CoolQMap map[string]interface{}

func NewCoolQ(url, token string) *CoolQ {
	coolq := &CoolQ{
		URL:   url,
		Token: token,
	}
	return coolq
}
