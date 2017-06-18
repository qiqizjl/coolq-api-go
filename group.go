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

//GroupList 群列表
type GroupList []struct {
	GroupID   int    `mapstructure:"group_id"`
	GroupName string `mapstructure:"group_name"`
}

//SendGroupMsg 发送群消息
func (coolq *CoolQ) SendGroupMsg(GroupID int, Message interface{}, IsRaw bool) error {
	// send_private_msg
	_, err := coolq.httpPOST("/send_group_msg", Map{
		"group_id": GroupID,
		"message":  Message,
		"is_raw":   IsRaw,
	})
	return err
}

//GetGroupMemberList 获得群成员列表
func (coolq *CoolQ) GetGroupMemberList(GroupID int) (*GroupMmeberList, error) {
	info, err := coolq.httpPOST("/get_group_member_list", Map{
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
	info, err := coolq.httpPOST("/get_group_member_info", Map{
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
	_, err := coolq.httpPOST("/set_group_kick", Map{
		"group_id":           GroupID,
		"user_id":            UserID,
		"reject_add_request": AddRequest,
	})
	return err
}

//SetGroupUserBan 禁言某个用户
func (coolq *CoolQ) SetGroupUserBan(GroupID, UserID, Time int) error {
	_, err := coolq.httpPOST("/set_group_ban", Map{
		"group_id": GroupID,
		"user_id":  UserID,
		"duration": Time,
	})
	return err
}

//SetGroupUserAnonymousBan 禁言某个匿名
func (coolq *CoolQ) SetGroupUserAnonymousBan(GroupID int, Flag string, Time int) error {
	_, err := coolq.httpPOST("/set_group_anonymous_ban", Map{
		"group_id": GroupID,
		"flag":     Flag,
		"duration": Time,
	})
	return err
}

//SetGroupBan 全体禁言
func (coolq *CoolQ) SetGroupBan(GroupID int, Enable bool) error {
	_, err := coolq.httpPOST("/set_group_whole_ban", Map{
		"group_id": GroupID,
		"enable":   Enable,
	})
	return err
}

//SetGroupAdmin 设置管理员
func (coolq *CoolQ) SetGroupAdmin(GroupID, UserID int, Enable bool) error {
	_, err := coolq.httpPOST("/set_group_admin", Map{
		"group_id": GroupID,
		"user_id":  UserID,
		"enable":   Enable,
	})
	return err
}

//SetGroupAnonymous 设置是否可匿名
func (coolq *CoolQ) SetGroupAnonymous(GroupID int, Enable bool) error {
	_, err := coolq.httpPOST("/set_group_anonymous", Map{
		"group_id": GroupID,
		"enable":   Enable,
	})
	return err
}

//SetGroupCard 设置群名片
func (coolq *CoolQ) SetGroupCard(GroupID, UserID int, Card string) error {
	_, err := coolq.httpPOST("/set_group_card", Map{
		"group_id": GroupID,
		"user_id":  UserID,
		"card":     Card,
	})
	return err
}

//SetGroupTitle 设置群头衔
func (coolq *CoolQ) SetGroupTitle(GroupID, UserID int, Title string, Duration int) error {
	_, err := coolq.httpPOST("/set_group_special_title", Map{
		"group_id":      GroupID,
		"user_id":       UserID,
		"special_title": Title,
		"duration":      Duration,
	})
	return err
}

//LeaveGroup 离群
func (coolq *CoolQ) LeaveGroup(GroupID int, IsDismiss bool) error {
	_, err := coolq.httpPOST("/set_group_leave", Map{
		"group_id":   GroupID,
		"is_dismiss": IsDismiss,
	})
	return err
}

//SetGroupRequest 处理群请求
func (coolq *CoolQ) SetGroupRequest(Flag, Type string, Approve bool, Reason string) error {
	_, err := coolq.httpPOST("/set_group_add_request", Map{
		"flag":    Flag,
		"type":    Type,
		"approve": Approve,
		"reason":  Reason,
	})
	return err
}

//GetGroupList 获得群组列表
func (coolq *CoolQ) GetGroupList() (*GroupList, error) {
	info, err := coolq.httpPOST("/get_group_list", nil)
	// GroupList
	if err != nil {
		return nil, err
	}
	var groupList GroupList
	err = mapstructure.Decode(info, &groupList)
	if err != nil {
		return nil, err
	}
	return &groupList, nil
}
