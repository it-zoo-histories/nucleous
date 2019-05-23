package payloads

/*TokenPayload - пакет авторизационный по токену*/
type TokenPayload struct {
	Token string `json:"access_token"`
}
