package payloads

import (
	"errors"
	"regexp"
)

/*LoginPayload - авторизационный пакет*/
type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`

	ValidateUsername bool
	ValidatePassword bool
}

/*Validate - валидация полей*/
func (pay *LoginPayload) Validate() *LoginPayload {
	if err := pay.validateUserName(); err != nil {
		pay.ValidateUsername = false
	} else {
		pay.ValidateUsername = true
	}
	if err2 := pay.validatePassword(); err2 != nil {
		pay.ValidatePassword = false
	} else {
		pay.ValidatePassword = true
	}
	return pay
}

func (pay *LoginPayload) validateUserName() error {
	match, _ := regexp.MatchString(usernamePattern, pay.Username)
	if match {
		return nil
	} else {
		return errors.New("unsupporting username")
	}
}

func (pay *LoginPayload) validatePassword() error {
	match, _ := regexp.MatchString(passwordPattern, pay.Password)
	if match {
		return nil
	} else {
		return errors.New("unsupporting password")
	}
}
