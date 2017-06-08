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
