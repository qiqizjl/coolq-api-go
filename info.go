package coolq

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type CoolQGetLogin struct {
	UserID   int64  `mapstructure:"user_id"`
	Nickname string `mapstructure:"nickname"`
}

type CoolQCookies struct {
	Cookies string `mapstructure:"cookies"`
}

type CoolQToken struct {
	Token int64
}

type StrangerInfo struct {
	UserID   int    `mapstructure:"user_id"`
	Nickname string `mapstructure:"nickname"`
	Sex      string `mapstructure:"sex"`
	Age      int    `mapstructure:"age"`
}

func (coolq *CoolQ) GetLoginInfo() (error, *CoolQGetLogin) {
	info, err := coolq.httpPOST("/get_login_info", nil)
	if err != nil {
		return err, nil
	}
	var loginInfo CoolQGetLogin
	err = mapstructure.Decode(info, &loginInfo)
	if err != nil {
		return err, nil
	}
	return nil, &loginInfo
}

func (coolq *CoolQ) GetCookies() (error, *CoolQCookies) {
	info, err := coolq.httpPOST("/get_cookies", nil)
	if err != nil {
		return err, nil
	}
	var cookieInfo CoolQCookies
	err = mapstructure.Decode(info, &cookieInfo)
	if err != nil {
		return err, nil
	}
	return nil, &cookieInfo
}

func (coolq *CoolQ) GetCsrfToken() (error, *CoolQToken) {
	info, err := coolq.httpPOST("/get_csrf_token", nil)
	if err != nil {
		return err, nil
	}
	var tokenInfo CoolQToken
	err = mapstructure.Decode(info, &tokenInfo)
	fmt.Println(tokenInfo)
	if err != nil {
		return err, nil
	}
	return nil, &tokenInfo
}

func (coolq *CoolQ) GetStrangerInfo(UserID int, NoCache bool) (*StrangerInfo, error) {
	info, err := coolq.httpPOST("/get_stranger_info", CoolQMap{
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
