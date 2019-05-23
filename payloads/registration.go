package payloads

import (
	"errors"
	"regexp"
)

/*RegistrationPayload - регистрационный пакет*/
type RegistrationPayload struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`

	ValidateUsername bool
	ValidatePassword bool
	ValidateEmail    bool
}

/*Validate - валидация полей*/
func (pay *RegistrationPayload) Validate() *RegistrationPayload {
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

	if err3 := pay.validateEmail(); err3 != nil {
		pay.ValidateEmail = false
	} else {
		pay.ValidateEmail = true
	}
	return pay
}

func (pay *RegistrationPayload) validateUserName() error {
	match, _ := regexp.MatchString(usernamePattern, pay.UserName)
	if match {
		return nil
	}
	return errors.New("unsupporting username")
}

func (pay *RegistrationPayload) validatePassword() error {
	match, _ := regexp.MatchString(passwordPattern, pay.Password)
	if match {
		return nil
	}
	return errors.New("unsupporting password")

}

func (pay *RegistrationPayload) validateEmail() error {
	match, _ := regexp.MatchString(emailPattern, pay.Email)
	if match {
		return nil
	}
	return errors.New("unsupporting email")

}
