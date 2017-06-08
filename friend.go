package coolq

func (coolq *CoolQ) SendPriviateMsg(UserID int, Message string, IsRaw bool) error {
	// send_private_msg
	_, err := coolq.httpPOST("/send_private_msg", Map{
		"user_id": UserID,
		"message": Message,
		"is_raw":  IsRaw,
	})
	return err
}

//SendLike 点赞
func (coolq *CoolQ) SendLike(UserID, Number int) error {
	_, err := coolq.httpPOST("/send_like", Map{
		"user_id": UserID,
		"times":   Number,
	})
	return err
}

func (coolq *CoolQ) SetFriendRequest(Flag, Type string, Approve bool, Reason string) error {
	_, err := coolq.httpPOST("/set_friend_add_request", Map{
		"flag":    Flag,
		"type":    Type,
		"approve": Approve,
		"reason":  Reason,
	})
	return err
}
