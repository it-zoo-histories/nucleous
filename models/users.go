package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

/*User - структура пользователя*/
type User struct {
	ID          bson.ObjectId `json:"user_id" bson:"_id"`
	Name        string        `json:"user_name" bson:"name"`
	Email       string        `json:"user_email" bson:"email"`
	Created     time.Time     `json:"user_created" bson:"created"`
	Updated     []time.Time   `json:"user_updates" bson:"updates"`
	Verificated bool          `json:"user_verificated" bson:"verificated"`
}
