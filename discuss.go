package coolq

//SendDiscussMsg 发送聊天室消息
func (coolq *CoolQ) SendDiscussMsg(DiscussID int, Message interface{}, IsRaw bool) error {
	_, err := coolq.httpPOST("/send_discuss_msg", Map{
		"discuss_id": DiscussID,
		"message":    Message,
		"is_raw":     IsRaw,
	})
	return err
}

//LeaveDiscuss 离开聊天室
func (coolq *CoolQ) LeaveDiscuss(DiscussID int) error {
	_, err := coolq.httpPOST("/set_discuss_leave", Map{
		"discuss_id": DiscussID,
	})
	return err
}
