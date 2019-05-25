package payloads

/*CreateUserPayload - пакет создания нового пользователя*/
type CreateUserPayload struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
