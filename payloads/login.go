package payloads

/*LoginPayload - авторизационный пакет*/
type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
