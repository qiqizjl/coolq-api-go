package coolq

//SendPriviateMsg 发送私聊
func (coolq *CoolQ) SendPriviateMsg(UserID int, Message interface{}, IsRaw bool) error {
	// send_private_msg
	_, err := coolq.httpPOST("/send_private_msg", Map{
		"user_id": UserID,
		"message": Message,
		"is_raw":  IsRaw,
	})
	return err
}

//SendPriviateMoreMsg 发送多个人消息
func (coolq *CoolQ) SendPriviateMoreMsg(UserIDs []int, Message interface{}, IsRaw bool) (bool, []error) {
	if len(UserIDs) <= 0 {
		return false, nil
	}
	errReturn := make([]error, 0)
	haveError := false
	for _, UserID := range UserIDs {
		// errReturn = append(errReturn, coolq.SendPriviateMsg(UserID, Message, IsRaw))
		returns := coolq.SendPriviateMsg(UserID, Message, IsRaw)
		if returns != nil {
			haveError = true
		}
		errReturn = append(errReturn, returns)
	}
	return haveError, errReturn
}

//SendLike 点赞
func (coolq *CoolQ) SendLike(UserID, Number int) error {
	_, err := coolq.httpPOST("/send_like", Map{
		"user_id": UserID,
		"times":   Number,
	})
	return err
}

//SetFriendRequest 处理好友请求
func (coolq *CoolQ) SetFriendRequest(Flag, Type string, Approve bool, Remark string) error {
	_, err := coolq.httpPOST("/set_friend_add_request", Map{
		"flag":    Flag,
		"approve": Approve,
		"remark":  Remark,
	})
	return err
}
