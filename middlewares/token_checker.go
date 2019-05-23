package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"nucleous/configuration"
	"nucleous/dao"
	"nucleous/enhancer"
	"nucleous/models"
	"nucleous/security"
	"strings"
)

/*JWTChecker - проверка токена в запросе*/
type JWTChecker struct {
	JWTMiddleWare *security.ECDSAMiddle
	// TokenDao      *dao.TokenDAO
	EResponser *enhancer.Responser
}

const (
	accesTokenName = "Authorization"
)

/*New - создание нового объекта прослойки*/
func (checker *JWTChecker) New(tokenDao *dao.TokenDAO, userDao *dao.UserDAO, jwtconf *configuration.JWTPay) *JWTChecker {
	checker.JWTMiddleWare = &security.ECDSAMiddle{}
	checker.JWTMiddleWare = checker.JWTMiddleWare.New(tokenDao, userDao, jwtconf)
	checker.EResponser = &enhancer.Responser{}
	return checker
}

/*JWTMiddleware - прослойка проверки токена в запросе*/
func (checker *JWTChecker) JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get(accesTokenName)
		tokenValue := strings.Split(tokenHeader, " ")[1]

		token, err := checker.findTokenInBD(tokenValue)

		fmt.Println(token)
		if err != nil {
			checker.EResponser.ResponseWithError(w, r, http.StatusUnauthorized, map[string]string{
				"status":  "401",
				"context": "nucleous.JWTMiddleware",
				"code":    "token not exist",
			})
		}

		err2 := checker.checkValidationToken(token)

		if err2 != nil {
			checker.EResponser.ResponseWithError(w, r, http.StatusUnauthorized, map[string]string{
				"status":  "403",
				"context": "nucleous.JWTMiddleware",
				"code":    "invalid token",
			})
		}
		next(w, r)
	})
}

func (checker *JWTChecker) findTokenInBD(tokenValue string) (*models.Token, error) {
	token, err := checker.JWTMiddleWare.TokenDao.CheckToken(tokenValue)
	return token, err
}

func (checker *JWTChecker) checkValidationToken(token *models.Token) error {
	result, username, err := checker.JWTMiddleWare.JWTEcdsa.ValidateToken(token.Value)

	log.Println("Check token validation for user: ", username)

	if !result {
		// TODO: add remove token from database
	}

	return err
}
