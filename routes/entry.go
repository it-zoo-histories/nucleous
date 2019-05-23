package routes

import (
	"net/http"
	"nucleous/dao"
	"nucleous/enhancer"
	"nucleous/middlewares"

	"github.com/gorilla/mux"
)

/*EntryPointRoute - точка входа*/
type EntryPointRoute struct {
	TokenDao   *dao.TokenDAO
	UserDao    *dao.UserDAO
	EResponser *enhancer.Responser
}

func (route *EntryPointRoute) methodsNotAllowed(w http.ResponseWriter, r *http.Request) {
	route.EResponser.ResponseWithError(w, r, http.StatusMethodNotAllowed, map[string]string{
		"status": "error request!",
	})
}

func (route *EntryPointRoute) methodsNotFoundRequest(w http.ResponseWriter, r *http.Request) {
	route.EResponser.ResponseWithError(w, r, http.StatusNotFound, map[string]string{
		"status": "not found!",
	})
}

func (route *EntryPointRoute) configureErrorRoutes(router *mux.Router) *mux.Router {
	router.NotFoundHandler = http.HandlerFunc(route.methodsNotFoundRequest)
	router.MethodNotAllowedHandler = http.HandlerFunc(route.methodsNotAllowed)
	return router
}

/*SettingRouter - конфигурирование роутера*/
func (route *EntryPointRoute) SettingRouter(
	tokenD *dao.TokenDAO,
	userD *dao.UserDAO,
	jwtMiddle *middlewares.JWTChecker,
) *mux.Router {

	route.TokenDao = tokenD
	route.UserDao = userD
	route.EResponser = &enhancer.Responser{}

	router := mux.NewRouter()

	authRoute := &AuthRoute{}
	authRoute = authRoute.InitRoute(tokenD, userD, jwtMiddle)

	userRoute := &UserRoute{}
	userRoute = userRoute.InitRoute(tokenD, userD, jwtMiddle)

	router = route.configureErrorRoutes(router)

	return userRoute.RoutesSetting(authRoute.RoutesSetting(router))
}
