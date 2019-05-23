package configuration

/*DatabaseConfiguration - конфигурация бд*/
type DatabaseConfiguration struct {
	Address  string `json:"address"`
	UserName string `json:"username"`
	Password string `json:"password"`
}
