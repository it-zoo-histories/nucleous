package configuration

import "time"

/*JWTPay - конфигурация генерации ключей*/
type JWTPay struct {
	NameServer              string        `json:"nameserver"`
	KeyName                 string        `json:"keyname"`
	PrivateKeyPath          string        `json:"private_key"`
	PublicKeyPath           string        `json:"public_key"`
	AmountHoursExpiredToken time.Duration `json:"token_time_expired"`
	LengthKey               string        `json:"length_key"`
}
