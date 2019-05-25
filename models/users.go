package models

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

/*User - структура пользователя*/
type User struct {
	ID          bson.ObjectId `json:"user_id"`
	Name        string        `json:"user_name"`
	Avatar      string        `json:"avatar"`
	Email       string        `json:"user_email"`
	Password    []byte        `json:"user_password"`
	Created     time.Time     `json:"user_created"`
	Updated     []time.Time   `json:"user_updates"`
	Verificated bool          `json:"user_verificated"`
}

/*HashAndSalt - солирование пароля пользователя*/
func (usr *User) HashAndSalt(pwd string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	usr.Password = hash
	return nil
}

/*CheckPwdHash - проверка пароля пользователя по хешу*/
func (usr *User) CheckPwdHash(password string) error {
	return bcrypt.CompareHashAndPassword(usr.Password, []byte(password))
}
