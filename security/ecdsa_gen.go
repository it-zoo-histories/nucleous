package security

import (
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
func (middleware *ECDSAMiddle) Login() error {}

/*Registration - регистрация*/
func (middleware *ECDSAMiddle) Registration() error {}

/*Logout - выход из аккаунта*/
func (middleware *ECDSAMiddle) Logout() error {}

/*New - создание новой прослойки*/
func (middleware *ECDSAMiddle) New(tokensDao *dao.TokenDAO, usersDao *dao.UserDAO) *ECDSAMiddle {
	jwtMiddle, err := gojwt.NewGojwtECDSA("atom_nucleous", "Access_key", "./keys/priv_key.pem", "./keys/pub_key.pem", "384", 1)
}
