package security

import (
	"fmt"
	"nucleous/configuration"
	"nucleous/dao"

	"github.com/jenazads/gojwt"
)

/*ECDSAMiddle - */
type ECDSAMiddle struct {
	TokenDao *dao.TokenDAO
	UserDao  *dao.UserDAO
	JWTEcdsa *gojwt.Gojwt
}

/*Login - авторизация*/
func (middleware *ECDSAMiddle) Login() error {
	fmt.Println("test")
	return nil
}

/*Registration - регистрация*/
func (middleware *ECDSAMiddle) Registration() error {
	return nil
}

/*Logout - выход из аккаунта*/
func (middleware *ECDSAMiddle) Logout() error {
	return nil
}

/*New - создание новой прослойки*/
func (middleware *ECDSAMiddle) New(tokensDao *dao.TokenDAO, usersDao *dao.UserDAO, config *configuration.JWTPay) *ECDSAMiddle {
	jwtMiddle, err := gojwt.NewGojwtECDSA(config.NameServer, config.KeyName, config.PrivateKeyPath, config.PublicKeyPath, config.LengthKey, config.AmountHoursExpiredToken)
	middleware.JWTEcdsa = jwtMiddle
	return nil
}
