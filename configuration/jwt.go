package configuration

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

/*JWTPay - конфигурация генерации ключей*/
type JWTPay struct {
	NameServer              string        `json:"nameserver"`
	KeyName                 string        `json:"keyname"`
	PrivateKeyPath          string        `json:"private_key"`
	PublicKeyPath           string        `json:"public_key"`
	AmountHoursExpiredToken time.Duration `json:"token_time_expired"`
	LengthKey               string        `json:"length_key"`
}

/*Parse - парсинг аргументов из файликов*/
func (obj *JWTPay) Parse(pathname string) (*JWTPay, error) {
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
