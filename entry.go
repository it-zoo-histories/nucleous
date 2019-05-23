package main

import (
	"fmt"
	"nucleous/configuration"
)

const (
	nameService = "[NUCLEOUS]"
)

var (
	mapConfigs     []string
	databaseConfig *configuration.DatabaseConfiguration
	initConfig     *configuration.InitPackage
	jwtConfig      *configuration.JWTPay
)

func parsedParams() {
	databaseConfig, _ = databaseConfig.Parse(mapConfigs[0])
	initConfig, _ = initConfig.Parse(mapConfigs[1])
	jwtConfig, _ = jwtConfig.Parse(mapConfigs[2])
}

func init() {
	databaseConfig = &configuration.DatabaseConfiguration{}
	initConfig = &configuration.InitPackage{}
	jwtConfig = &configuration.JWTPay{}

	mapConfigs = []string{
		"./settings/database.json",
		"./settings/init.json",
		"./settings/jwt.json",
	}
	parsedParams()
}

func main() {
	fmt.Println(nameService + " was started")
}
