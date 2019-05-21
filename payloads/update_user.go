package payloads

/*UserUpdatePayload - */
type UserUpdatePayload struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (pay *UserUpdatePayload) Validate() error {
	return nil
}

func (pay *UserUpdatePayload) validateUserName() error {
	return nil
}

func (pay *UserUpdatePayload) validatePassword() error {
	return nil
}

func (pay *UserUpdatePayload) validateEmail() error {
	return nil
}
