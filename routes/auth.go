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
			"context": "AuthRoute.Logout",
		})
	}
}

/*logoutFromAccount - выход из аккаунта*/
func (route *AuthRoute) logoutFromAccount(w http.ResponseWriter, r *http.Request) {
	var payload payloads.LogoutPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "AuthRoute.Logout",
		})
	}
}

/*regsitrationAccount - регистрация нового аккаунта*/
func (route *AuthRoute) regsitrationAccount(w http.ResponseWriter, r *http.Request) {
	var payload payloads.RegistrationPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "AuthRoute.Registration",
		})
	}
}

/*autheticationByToken - авторизация и прокидка по токену*/
func (route *AuthRoute) autheticationByToken(w http.ResponseWriter, r *http.Request) {
	var payload payloads.TokenPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "AuthRoute.AuthenticationByToken",
		})
	}
}

/*RoutesSetting - конфигурация роутера для маршрутов авторизации\регистрации*/
func (route *AuthRoute) RoutesSetting(router *mux.Router) *mux.Router {
	router.HandleFunc(authenticationByTokenAddress, route.jwtMiddle.JWTMiddleware(route.resendMiddle.SendNext(route.autheticationByToken))).Methods("POST")
	router.HandleFunc(authenticationByCredentials, route.autheticationByCredentials).Methods("POST")
	router.HandleFunc(logoutFromAccount, route.jwtMiddle.JWTMiddleware(route.logoutFromAccount)).Methods("POST")
	router.HandleFunc(registrationAccount, route.regsitrationAccount).Methods("POST")
	return router
}
