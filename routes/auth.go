package routes

import (
	"net/http"
	"nucleous/dao"

	"github.com/gorilla/mux"
)

/*AuthRoute - авторизационный маршрут (для регистрации, авторизации)*/
type AuthRoute struct {
	TokenDao *dao.TokenDAO
	UserDao  *dao.UserDAO
}

const (
	authenticationByTokenAddress = "/auth/token"
	authenticationByCredentials  = "/auth/cred"
	logoutFromAccount            = "/auth/logout"
	registrationAccount          = "/auth/registration"
)

/*InitRoute - инициализация роутера*/
func (route *AuthRoute) InitRoute(tokens *dao.TokenDAO, users *dao.UserDAO) *AuthRoute {
	route.TokenDao = tokens
	route.UserDao = users
	return route
}

/*autheticationByCredentials - вход по данным логина и пароля*/
func (route *AuthRoute) autheticationByCredentials(w http.ResponseWriter, r *http.Request) {

}

/*logoutFromAccount - выход из аккаунта*/
func (route *AuthRoute) logoutFromAccount(w http.ResponseWriter, r *http.Request) {

}

/*regsitrationAccount - регистрация нового аккаунта*/
func (route *AuthRoute) regsitrationAccount(w http.ResponseWriter, r *http.Request) {

}

/*autheticationByToken - авторизация по токену*/
func (route *AuthRoute) autheticationByToken(w http.ResponseWriter, r *http.Request) {

}

/*RoutesSetting - конфигурация роутера для маршрутов авторизации\регистрации*/
func (route *AuthRoute) RoutesSetting(router *mux.Router) *mux.Router {
	router.HandleFunc(authenticationByTokenAddress, route.autheticationByToken).Methods("POST")
	router.HandleFunc(authenticationByCredentials, route.autheticationByCredentials).Methods("POST")
	router.HandleFunc(logoutFromAccount, route.logoutFromAccount).Methods("POST")
	router.HandleFunc(registrationAccount, route.regsitrationAccount).Methods("POST")
	return router
}
