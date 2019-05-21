package configuration

import (
	"nucleous/dao"
)

/*JWTPay - конфигурация генерации ключей*/
type JWTPay struct {
	NameServer              string         `json:"nameserver"`
	KeyName                 string         `json:"keyname"`
	PrivateKeyPath          string         `json:"private_key"`
	PublicKeyPath           string         `json:"public_key"`
	AmountHoursExpiredToken int            `json:"token_time_expired"`
	LengthKey               int            `json:"length_key"`
	TokensDao               *dao.TokensDao `json:"-"`
	UsersDao                *dao.UserDAO   `json:"-"`
}
