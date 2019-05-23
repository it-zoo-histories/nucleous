package routes

import (
	"net/http"
	"nucleous/dao"

	"github.com/gorilla/mux"
)

/*UserRoute - маршрут для изменения данных пользователя и его удаления*/
type UserRoute struct {
	TokenDao *dao.TokenDAO
	UserDao  *dao.UserDAO
}

const (
	updateUserSettings = "/user"
	removeUserByID     = "/user"
)

func (route *UserRoute) updateUserSettings(w http.ResponseWriter, r *http.Request) {

}

func (route *UserRoute) removeUserByID(w http.ResponseWriter, r *http.Request) {

}

/*InitRoute - инициализация роутера*/
func (route *UserRoute) InitRoute(tokens *dao.TokenDAO, users *dao.UserDAO) *UserRoute {
	route.TokenDao = tokens
	route.UserDao = users
	return route
}

/*RoutesSetting - конфигурация роутера для маршрутов авторизации\регистрации*/
func (route *UserRoute) RoutesSetting(router *mux.Router) *mux.Router {
	router.HandleFunc(updateUserSettings, route.removeUserByID).Methods("UPDATE")
	router.HandleFunc(removeUserByID, route.removeUserByID).Methods("DELETE")
	return router
}
