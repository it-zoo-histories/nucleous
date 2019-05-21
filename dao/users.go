package dao

import (
	"nucleous/models"
	"time"

	"gopkg.in/mgo.v2"
)

type UserDAO struct {
	Database *mgo.Database
}

const (
	UsersCollection = "users"
)

func (dao *UserDAO) New(database *mgo.Database) *UserDAO {
	dao.Database = database
	return dao
}

/*CreateNewUser - создание нового пользователя*/
func (dao *UserDAO) CreateNewUser(user *models.User) (err error) {
	user.Created = time.Now()
	user.Verificated = false

	err = dao.Database.C(UsersCollection).Insert(&user)
	return
}

/*UpdateUser - обновление данных пользователя по его идентификатору*/
func (dao *UserDAO) UpdateUser(userid string, newpayload *models.User) (updated *models.User, err error) {

}
