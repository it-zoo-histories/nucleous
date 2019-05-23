package configuration

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

/*DatabaseConfiguration - конфигурация бд*/
type DatabaseConfiguration struct {
	Address  string `json:"address"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

/*Parse - парсинг аргументов из файликов*/
func (obj *DatabaseConfiguration) Parse(pathname string) (*DatabaseConfiguration, error) {
	jsonFile, err := os.Open(pathname)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)

	err2 := json.Unmarshal(bytes, obj)

	if err2 != nil {
		return nil, err2
	}

	return obj, nil
}
