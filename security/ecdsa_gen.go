package security

import (
	"nucleous/configuration"
	"nucleous/dao"
	"nucleous/models"
	"nucleous/payloads"
	"time"

	"github.com/jenazads/gojwt"
)

/*ECDSAMiddle - */
type ECDSAMiddle struct {
	TokenDao *dao.TokenDAO
	UserDao  *dao.UserDAO
	JWTEcdsa *gojwt.Gojwt
}

/*CreateNewToken - авторизация*/
func (middleware *ECDSAMiddle) CreateNewToken(obj *payloads.LoginPayload, userFromDb *models.User) (*models.Token, error) {
	// fmt.Println("test")
	tokenValue, err := middleware.JWTEcdsa.CreateToken(obj.Username)
	if err != nil {
		return nil, err
	}

	tokenUser := &models.Token{}
	tokenUser.Value = tokenValue
	tokenUser.UserID = userFromDb.ID
	tokenUser.Created = time.Now()
	tokenUser.ExpireTo = time.Hour * 2

	token, err2 := middleware.TokenDao.CreateNewToken(tokenUser)
	if err2 != nil {
		return nil, err2
	}
	return token, nil
}

/*New - создание новой прослойки*/
func (middleware *ECDSAMiddle) New(tokensDao *dao.TokenDAO, usersDao *dao.UserDAO, config *configuration.JWTPay) *ECDSAMiddle {
	jwtMiddle, _ := gojwt.NewGojwtECDSA(config.NameServer, config.KeyName, config.PrivateKeyPath, config.PublicKeyPath, config.LengthKey, config.AmountHoursExpiredToken)
	middleware.JWTEcdsa = jwtMiddle
	middleware.TokenDao = tokensDao
	middleware.UserDao = usersDao
	return middleware
}

/*tokenValidation - проверка токена*/
func (middleware *ECDSAMiddle) tokenValidation(token string) (bool, error) {
	result, _, err := middleware.JWTEcdsa.ValidateToken(token)
	return result, err
}
