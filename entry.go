package main

import (
	"fmt"
	"log"
	"net/http"
	"nucleous/configuration"
	"nucleous/dao"
	"nucleous/middlewares"
	"nucleous/routes"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	nameService = "[NUCLEOUS]"
)

var (
	mapConfigs     []string
	databaseConfig *configuration.DatabaseConfiguration
	initConfig     *configuration.InitPackage
	jwtConfig      *configuration.JWTPay
	jwtMiddle      *middlewares.JWTChecker
	entryRoute     *routes.EntryPointRoute
	tokenDao       *dao.TokenDAO
	userDao        *dao.UserDAO
	database       *dao.Database
	router         *mux.Router
)

func parsedParams() {
	databaseConfig, _ = databaseConfig.Parse(mapConfigs[0])
	initConfig, _ = initConfig.Parse(mapConfigs[1])
	jwtConfig, _ = jwtConfig.Parse(mapConfigs[2])
}

func configureDaos() {
	tokenDao = &dao.TokenDAO{}
	userDao = &dao.UserDAO{}
	tokenDao = tokenDao.New(database)
	userDao = userDao.New(database)
}

func databasePreparing() {
	err := database.ConnectionUser(databaseConfig)
	if err != nil {
		log.Println(nameService+"Error connected to database: ", err.Error())
	}

	err2 := database.ConnectToBD()
	if err2 != nil {
		log.Println(nameService+"error connect to database: ", err2.Error())
	}
}

func configureJWT() {
	jwtMiddle := &middlewares.JWTChecker{}
	jwtMiddle = jwtMiddle.New(tokenDao, userDao, jwtConfig)
}

func init() {
	databaseConfig = &configuration.DatabaseConfiguration{}
	initConfig = &configuration.InitPackage{}
	jwtConfig = &configuration.JWTPay{}
	entryRoute = &routes.EntryPointRoute{}
	database = &dao.Database{}

	entryRoute := routes.EntryPointRoute{}
	mapConfigs = []string{
		"./settings/database.json",
		"./settings/init.json",
		"./settings/jwt.json",
	}
	parsedParams()
	databasePreparing()
	configureDaos()
	configureJWT()

	router = entryRoute.SettingRouter(tokenDao, userDao, jwtMiddle)

}

func routerForListen() http.Handler {
	return applyCorsPolicy()
}

func applyCorsPolicy() http.Handler {
	handler := cors.New(
		cors.Options{
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders: []string{"Content-Type", "Accept", "Authorization", "Access-Control-Allow-Origin"},
			AllowedOrigins: []string{"*"},
			Debug:          true,
		},
	).Handler(router)

	return handler
}

func main() {
	if err := http.ListenAndServe(initConfig.ServerAddress+":"+strconv.Itoa(initConfig.ServerPort), routerForListen()); err != nil {
		log.Fatalln(nameService+"Error starting api! err code: ", err.Error())
	}
	fmt.Println(nameService + " was started")
}
