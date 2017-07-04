package coolq

import (
	"github.com/mitchellh/mapstructure"
)

//LoginInfo 登录信息
type LoginInfo struct {
	UserID   int64  `mapstructure:"user_id"`
	Nickname string `mapstructure:"nickname"`
}

//Cookies Cookies信息
type Cookies struct {
	Cookies string `mapstructure:"cookies"`
}

//Token  Token信息
type Token struct {
	Token int64
}

//StrangerInfo 陌生人信息
type StrangerInfo struct {
	UserID   int    `mapstructure:"user_id"`
	Nickname string `mapstructure:"nickname"`
	Sex      string `mapstructure:"sex"`
	Age      int    `mapstructure:"age"`
}

//SystemInfo 插件信息
type SystemInfo struct {
	CoolqEdition  string `mapstructure:"coolq_edition"`
	PluginVersion string `mapstructure:"plugin_version"`
}

//GetLoginInfo 获得当前登录信息
func (coolq *CoolQ) GetLoginInfo() (*LoginInfo, error) {
	info, err := coolq.httpPOST("/get_login_info", nil)
	if err != nil {
		return nil, err
	}
	var loginInfo LoginInfo
	err = mapstructure.Decode(info, &loginInfo)
	if err != nil {
		return nil, err
	}
	return &loginInfo, nil
}

//GetCookies 获得登录Cookie
func (coolq *CoolQ) GetCookies() (*Cookies, error) {
	info, err := coolq.httpPOST("/get_cookies", nil)
	if err != nil {
		return nil, err
	}
	var cookieInfo Cookies
	err = mapstructure.Decode(info, &cookieInfo)
	if err != nil {
		return nil, err
	}
	return &cookieInfo, nil
}

//GetCsrfToken 获得Token
func (coolq *CoolQ) GetCsrfToken() (*Token, error) {
	info, err := coolq.httpPOST("/get_csrf_token", nil)
	if err != nil {
		return nil, err
	}
	var tokenInfo Token
	err = mapstructure.Decode(info, &tokenInfo)
	if err != nil {
		return nil, err
	}
	return &tokenInfo, nil
}

//GetStrangerInfo 获得陌生人信息
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

//GetSystemInfo 获得版本号信息
func (coolq *CoolQ) GetSystemInfo() (*SystemInfo, error) {
	info, err := coolq.httpPOST("/get_version_info", Map{})
	if err != nil {
		return nil, err
	}
	var systmInfo SystemInfo
	err = mapstructure.Decode(info, &systmInfo)
	if err != nil {
		return nil, err
	}
	return &systmInfo, nil
}
