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

/*UserRoute - маршрут для изменения данных пользователя и его удаления*/
type UserRoute struct {
	TokenDao   *dao.TokenDAO
	UserDao    *dao.UserDAO
	JWTMiddle  *middlewares.JWTChecker
	EResponser *enhancer.Responser
}

const (
	updateUserSettings = "/user"
	removeUserByID     = "/user/remove"
)

func (route *UserRoute) updateUserSettings(w http.ResponseWriter, r *http.Request) {
	var payload payloads.UserUpdatePayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "UserRoute.UpdateUserSettings",
		})
	}
}

func (route *UserRoute) removeUserByID(w http.ResponseWriter, r *http.Request) {
	var payload payloads.RemoveUserPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "UserRoute.RemoveUserByID",
		})
	}
}

/*InitRoute - инициализация роутера*/
func (route *UserRoute) InitRoute(tokens *dao.TokenDAO, users *dao.UserDAO, jwtMiddle *middlewares.JWTChecker) *UserRoute {
	route.TokenDao = tokens
	route.UserDao = users
	route.JWTMiddle = jwtMiddle
	return route
}

/*RoutesSetting - конфигурация роутера для маршрутов авторизации\регистрации*/
func (route *UserRoute) RoutesSetting(router *mux.Router) *mux.Router {
	router.HandleFunc(updateUserSettings, route.JWTMiddle.JWTMiddleware(route.removeUserByID)).Methods("UPDATE")
	router.HandleFunc(removeUserByID, route.JWTMiddle.JWTMiddleware(route.removeUserByID)).Methods("DELETE")
	return router
}
