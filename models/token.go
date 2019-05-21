package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

/*Token - сущнсоть токена */
type Token struct {
	ID       bson.ObjectId `json:"-" bson:"_id"`
	UserID   bson.ObjectId `json:"-" bson:"owner"`
	Value    string        `json:"value" bson:"value"`     // сгенерированное значение токена
	Created  time.Time     `json:"created" bson:"created"` // дата создания токена
	ExpireTo time.Time     `json:"-" bson:"validTo"`       // дата, до которой будет валиден токен
}
