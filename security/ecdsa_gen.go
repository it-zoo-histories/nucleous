package security

import (
	"fmt"
	"nucleous/configuration"
	"nucleous/dao"
	"nucleous/payloads"

	"github.com/jenazads/gojwt"
)

/*ECDSAMiddle - */
type ECDSAMiddle struct {
	TokenDao *dao.TokenDAO
	UserDao  *dao.UserDAO
	JWTEcdsa *gojwt.Gojwt
}

/*Login - авторизация*/
func (middleware *ECDSAMiddle) Login(obj *payloads.LoginPayload) error {
	fmt.Println("test")
	return nil
}

/*Registration - регистрация*/
func (middleware *ECDSAMiddle) Registration(obj *payloads.RegistrationPayload) error {
	return nil
}

/*Logout - выход из аккаунта*/
func (middleware *ECDSAMiddle) Logout(obj *payloads.LogoutPayload) error {
	return nil
}

/*New - создание новой прослойки*/
func (middleware *ECDSAMiddle) New(tokensDao *dao.TokenDAO, usersDao *dao.UserDAO, config *configuration.JWTPay) *ECDSAMiddle {
	jwtMiddle, _ := gojwt.NewGojwtECDSA(config.NameServer, config.KeyName, config.PrivateKeyPath, config.PublicKeyPath, config.LengthKey, config.AmountHoursExpiredToken)
	middleware.JWTEcdsa = jwtMiddle
	middleware.TokenDao = tokensDao
	middleware.UserDao = usersDao
	return nil
}

/*CheckEnterUserToDatabase - проверка пользователя в базе данных*/
func (middleware *ECDSAMiddle) CheckEnterUserToDatabase(obj payloads.LoginPayload) {

}

/*tokenValidation - проверка токена*/
func (middleware *ECDSAMiddle) tokenValidation(token string) (bool, error) {
	result, _, err := middleware.JWTEcdsa.ValidateToken(token)
	return result, err
}
