package coolq

import "github.com/mitchellh/mapstructure"

//GroupMemberInfo 用户信息
type GroupMemberInfo struct {
	GroupID         int    `mapstructure:"group_id"`
	UserID          int    `mapstructure:"user_id"`
	Nickname        string `mapstructure:"nickname"`
	Card            string `mapstructure:"card"`
	Sex             string `mapstructure:"sex"`
	Age             int    `mapstructure:"age"`
	Area            string `mapstructure:"area"`
	JoinTime        int    `mapstructure:"join_time"`
	LastSentTime    int    `mapstructure:"last_sent_time"`
	Level           string `mapstructure:"level"`
	Role            string `mapstructure:"role"`
	Unfriendly      bool   `mapstructure:"unfriendly"`
	Title           string `mapstructure:"title"`
	TitleExpireTime int    `mapstructure:"title_expire_time"`
	CardChangeable  bool   `mapstructure:"card_changeable"`
}

//GroupMmeberList 用户列表
type GroupMmeberList []GroupMemberInfo

//SendGroupMsg 发送群消息
func (coolq *CoolQ) SendGroupMsg(GourpID int, Message string, IsRaw bool) error {
	// send_private_msg
	_, err := coolq.httpPOST("/send_group_msg", CoolQMap{
		"group_id": GourpID,
		"message":  Message,
		"is_raw":   IsRaw,
	})
	return err
}

//GetGroupMemberList 获得群成员列表
func (coolq *CoolQ) GetGroupMemberList(GroupID int) (*GroupMmeberList, error) {
	info, err := coolq.httpPOST("/get_group_member_list", CoolQMap{
		"group_id": GroupID,
	})
	if err != nil {
		return nil, err
	}
	var memberList GroupMmeberList
	err = mapstructure.Decode(info, &memberList)
	if err != nil {
		return nil, err
	}
	return &memberList, nil
}

//GetGroupMemberInfo 获得群成员信息
func (coolq *CoolQ) GetGroupMemberInfo(GroupID, UserID int, NoCache bool) (*GroupMemberInfo, error) {
	info, err := coolq.httpPOST("/get_group_member_info", CoolQMap{
		"group_id": GroupID,
		"user_id":  UserID,
		"no_cache": NoCache,
	})
	if err != nil {
		return nil, err
	}
	var memberInfo GroupMemberInfo
	err = mapstructure.Decode(info, &memberInfo)
	if err != nil {
		return nil, err
	}
	return &memberInfo, nil
}

//SetGroupKick 群组踢人
func (coolq *CoolQ) SetGroupKick(GroupID, UserID int, AddRequest bool) error {
	_, err := coolq.httpPOST("/set_group_kick", CoolQMap{
		"group_id":           GroupID,
		"user_id":            UserID,
		"reject_add_request": AddRequest,
	})
	return err
}

func (coolq *CoolQ) SetGroupUserBan(GroupID, UserID, Time int) error {
	_, err := coolq.httpPOST("/set_group_ban", CoolQMap{
		"group_id": GroupID,
		"user_id":  UserID,
		"duration": Time,
	})
	return err
}

func (coolq *CoolQ) SetGroupUserAnonymousBan(GroupID int, Flag string, Time int) error {
	_, err := coolq.httpPOST("/set_group_anonymous_ban", CoolQMap{
		"group_id": GroupID,
		"flag":     Flag,
		"duration": Time,
	})
	return err
}

func (coolq *CoolQ) SetGroupBan(GroupID int, Enable bool) error {
	_, err := coolq.httpPOST("/set_group_whole_ban", CoolQMap{
		"group_id": GroupID,
		"enable":   Enable,
	})
	return err
}

func (coolq *CoolQ) SetGroupAdmin(GroupID, UserID int, Enable bool) error {
	_, err := coolq.httpPOST("/set_group_admin", CoolQMap{
		"group_id": GroupID,
		"user_id":  UserID,
		"enable":   Enable,
	})
	return err
}

func (coolq *CoolQ) SetGroupAnonymous(GroupID int, Enable bool) error {
	_, err := coolq.httpPOST("/set_group_anonymous", CoolQMap{
		"group_id": GroupID,
		"enable":   Enable,
	})
	return err
}

func (coolq *CoolQ) SetGroupCard(GroupID, UserID int, Card string) error {
	_, err := coolq.httpPOST("/set_group_card", CoolQMap{
		"group_id": GroupID,
		"user_id":  UserID,
		"card":     Card,
	})
	return err
}

func (coolq *CoolQ) SetGroupTitle(GroupID, UserID int, Title string, Duration int) error {
	_, err := coolq.httpPOST("/set_group_special_title", CoolQMap{
		"group_id":      GroupID,
		"user_id":       UserID,
		"special_title": Title,
		"duration":      Duration,
	})
	return err
}

func (coolq *CoolQ) LeaveGroup(GroupID int, IsDismiss bool) error {
	_, err := coolq.httpPOST("/set_group_leave", CoolQMap{
		"group_id":   GroupID,
		"is_dismiss": IsDismiss,
	})
	return err
}

func (coolq *CoolQ) SetGroupRequest(Flag, Type string, Approve bool, Reason string) error {
	_, err := coolq.httpPOST("/set_group_add_request", CoolQMap{
		"flag":    Flag,
		"type":    Type,
		"approve": Approve,
		"reason":  Reason,
	})
	return err
}
