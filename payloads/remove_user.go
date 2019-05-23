package payloads

/*RemoveUserPayload - удаление юзера по его идентификатору*/
type RemoveUserPayload struct {
	UserID string `json:"user_id"`
}
