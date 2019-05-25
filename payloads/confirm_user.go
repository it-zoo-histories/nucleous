package payloads

/*ConfirmCodeUserPayload - подвтерждающий пакет*/
type ConfirmCodeUserPayload struct {
	Code   int    `json:"code"`
	Userid string `json:"userid"`
}
