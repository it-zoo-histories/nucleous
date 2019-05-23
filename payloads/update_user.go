package payloads

import (
	"errors"
	"regexp"
)

/*UserUpdatePayload - */
type UserUpdatePayload struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`

	ValidateUsername bool
	ValidatePassword bool
	ValidateEmail    bool
}

const (
	usernamePattern = "([A-Za-z_0-9])\\w{4,}"
	passwordPattern = "^(?=.*\\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z]).{8,}$"
	emailPattern    = "^([A-Z|a-z|0-9](\\.|_){0,1})+[A-Z|a-z|0-9]\\@([A-Z|a-z|0-9])+((\\.){0,1}[A-Z|a-z|0-9]){2}\\.[a-z]{2,3}$"
)

/*Validate - валидация полей*/
func (pay *UserUpdatePayload) Validate() *UserUpdatePayload {
	if err := pay.validateUserName(); err != nil {
		pay.ValidateUsername = false
	} else {
		pay.ValidatePassword = true
	}
	if err2 := pay.validatePassword(); err2 != nil {
		pay.ValidatePassword = false
	} else {
		pay.ValidatePassword = true
	}
	if err3 := pay.validateEmail(); err3 != nil {
		pay.ValidateEmail = false
	} else {
		pay.ValidateEmail = true
	}
	return pay
}

func (pay *UserUpdatePayload) validateUserName() error {
	match, _ := regexp.MatchString(usernamePattern, pay.UserName)
	if match {
		return nil
	} else {
		return errors.New("unsupporting username")
	}
}

func (pay *UserUpdatePayload) validatePassword() error {
	match, _ := regexp.MatchString(passwordPattern, pay.Password)
	if match {
		return nil
	} else {
		return errors.New("unsupporting password")
	}
}

func (pay *UserUpdatePayload) validateEmail() error {
	match, _ := regexp.MatchString(emailPattern, pay.Email)
	if match {
		return nil
	} else {
		return errors.New("unsupporting email")
	}
}
