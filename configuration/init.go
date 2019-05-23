package configuration

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

/*InitPackage - инициализирующий пакет*/
type InitPackage struct {
	ServerAddress string `json:"server_address"`
	ServerPort    int    `json:"server_port"`
}

/*Parse - парсинг аргументов из файлика*/
func (obj *InitPackage) Parse(pathname string) (*InitPackage, error) {
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
