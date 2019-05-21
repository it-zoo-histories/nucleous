package payloads

/*RegistrationPayload - регистрационный пакет*/
type RegistrationPayload struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
