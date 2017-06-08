package coolq

import (
	"github.com/mitchellh/mapstructure"
)

type LoginInfo struct {
	UserID   int64  `mapstructure:"user_id"`
	Nickname string `mapstructure:"nickname"`
}

type Cookies struct {
	Cookies string `mapstructure:"cookies"`
}

type Token struct {
	Token int64
}

type StrangerInfo struct {
	UserID   int    `mapstructure:"user_id"`
	Nickname string `mapstructure:"nickname"`
	Sex      string `mapstructure:"sex"`
	Age      int    `mapstructure:"age"`
}

func (coolq *CoolQ) GetLoginInfo() (error, *LoginInfo) {
	info, err := coolq.httpPOST("/get_login_info", nil)
	if err != nil {
		return err, nil
	}
	var loginInfo LoginInfo
	err = mapstructure.Decode(info, &loginInfo)
	if err != nil {
		return err, nil
	}
	return nil, &loginInfo
}

func (coolq *CoolQ) GetCookies() (error, *Cookies) {
	info, err := coolq.httpPOST("/get_cookies", nil)
	if err != nil {
		return err, nil
	}
	var cookieInfo Cookies
	err = mapstructure.Decode(info, &cookieInfo)
	if err != nil {
		return err, nil
	}
	return nil, &cookieInfo
}

func (coolq *CoolQ) GetCsrfToken() (error, *Token) {
	info, err := coolq.httpPOST("/get_csrf_token", nil)
	if err != nil {
		return err, nil
	}
	var tokenInfo Token
	err = mapstructure.Decode(info, &tokenInfo)
	if err != nil {
		return err, nil
	}
	return nil, &tokenInfo
}

func (coolq *CoolQ) GetStrangerInfo(UserID int, NoCache bool) (*StrangerInfo, error) {
	info, err := coolq.httpPOST("/get_stranger_info", Map{
		"user_id":  UserID,
		"no_cache": NoCache,
	})
	if err != nil {
		return nil, err
	}
	var strangerInfo StrangerInfo
	err = mapstructure.Decode(info, &strangerInfo)
	if err != nil {
		return nil, err
	}
	return &strangerInfo, nil
}
