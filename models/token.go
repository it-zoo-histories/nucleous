package models

import (
	"time"
)

/*Token - сущнсоть токена */
type Token struct {
	ID       string        `json:"token_id`
	UserID   string        `json:"user_id"`
	Value    string        `json:"token_value"`  // сгенерированное значение токена
	Created  time.Time     `json:"created_time"` // дата создания токена
	ExpireTo time.Duration `json:"expired_time"` // дата, до которой будет валиден токен
}
