package routes

import (
	"encoding/json"
	"net/http"
	"nucleous/dao"
	"nucleous/enhancer"
	"nucleous/middlewares"
	"nucleous/payloads"

	"github.com/gorilla/mux"
)

/*AuthRoute - авторизационный маршрут (для регистрации, авторизации)*/
type AuthRoute struct {
	TokenDao     *dao.TokenDAO
	UserDao      *dao.UserDAO
	jwtMiddle    *middlewares.JWTChecker
	resendMiddle *middlewares.NextServe
	EResponser   *enhancer.Responser
}

const (
	authenticationByTokenAddress = "/auth/token"
	authenticationByCredentials  = "/auth/cred"
	logoutFromAccount            = "/auth/logout"
	registrationAccount          = "/auth/registration"
)

/*InitRoute - инициализация роутера*/
func (route *AuthRoute) InitRoute(
	tokens *dao.TokenDAO,
	users *dao.UserDAO,
	jwtMiddle *middlewares.JWTChecker,
) *AuthRoute {

	route.TokenDao = tokens
	route.UserDao = users
	route.EResponser = &enhancer.Responser{}
	route.jwtMiddle = jwtMiddle
	route.resendMiddle = &middlewares.NextServe{
		EResponse: &enhancer.Responser{},
	}

	return route
}

/*autheticationByCredentials - вход по данным логина и пароля*/
func (route *AuthRoute) autheticationByCredentials(w http.ResponseWriter, r *http.Request) {
	var payload payloads.LoginPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "AuthRoute.AuthenticationByCredentials",
			"code":    err.Error(),
		},
			"application/json",
		)
		return

	}

	usr, err2 := route.UserDao.GetUserByCredentials(&payload)
	if err2 != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "AuthRoute.AuthenticationByCredentials",
			"code":    err2.Error(),
		},
			"application/json",
		)

		return
	}

	tokenFromDb, err3 := route.TokenDao.GetTokenByUserID(usr.ID)
	if err3 != nil {
		newToken, err5 := route.jwtMiddle.JWTMiddleWare.CreateNewToken(&payload, usr)
		if err5 != nil {
			route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
				"status":  "error",
				"context": "AuthRoute.AuthenticationByCredentials",
				"code":    err5.Error(),
			},
				"application/json",
			)
			return
		}

		route.EResponser.ResponseWithJSON(w, r, http.StatusOK, newToken,
			"application/json",
		)
		return

	}

	jsonToken, err4 := json.Marshal(tokenFromDb)
	if err4 != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "AuthRoute.AuthenticationByCredentials",
			"code":    err4.Error(),
		},
			"application/json",
		)
		return
	}

	route.EResponser.ResponseWithJSON(w, r, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"context": "AuthRoute.AuthenticationByCredentials",
		"data":    string(jsonToken),
	},
		"application/json",
	)

	return

}

/*logoutFromAccount - выход из аккаунта*/
func (route *AuthRoute) logoutFromAccount(w http.ResponseWriter, r *http.Request) {
	var payload payloads.LogoutPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "AuthRoute.Logout",
			"code":    err.Error(),
		},
			"application/json",
		)
	}

	if err2 := route.TokenDao.RemoveTokenByID(payload.Token); err2 != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "AuthRoute.Logout",
			"code":    err2.Error(),
		},
			"application/json",
		)
	}
	route.EResponser.ResponseWithJSON(w, r, http.StatusOK, map[string]string{
		"status":  "success logout",
		"context": "AuthRoute.Logout",
	},
		"application/json",
	)
	return

}

// /*autheticationByToken - авторизация и прокидка по токену*/
// func (route *AuthRoute) autheticationByToken(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusUseProxy)
// }

/*RoutesSetting - конфигурация роутера для маршрутов авторизации\регистрации*/
func (route *AuthRoute) RoutesSetting(router *mux.Router) *mux.Router {
	router.HandleFunc(authenticationByTokenAddress, route.jwtMiddle.JWTMiddleware(route.resendMiddle.SendNext)).Methods("POST")
	router.HandleFunc(authenticationByCredentials, route.autheticationByCredentials).Methods("POST")
	router.HandleFunc(logoutFromAccount, route.jwtMiddle.JWTMiddleware(route.logoutFromAccount)).Methods("POST")
	return router
}
