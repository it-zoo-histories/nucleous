package dao

import (
	"errors"
	"nucleous/models"
	"time"

	"gopkg.in/mgo.v2"
)

type UserDAO struct {
	Database *mgo.Database
}

const (
	UsersCollection  = "users"
	ReservedTokenNil = "<nil>"
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
	iduser, errTranslate := bson.O
	if errTranslate != nil {
		return &models.User{}, errTranslate
	}

	var modelUser models.User

	err = dao.Database.C(UsersCollection).Find(bson.M{
		"_id": iduser,
	}).One(&modelUser)

	if err != nil {
		return &models.User{}, err
	}

	if errCheckUserVerification := dao.CheckVerificationUser(modelUser); errCheckUserVerification != nil {
		return errCheckUserVerification
	}

}

/*CheckVerificationUser - проверка подтверждённия аккаунта пользователя*/
func (dao *UserDAO) CheckVerificationUser(userModel models.User) error {
	if userModel.Verificated {
		return nil
	} else {
		return errors.New("user can not change user information")
	}
}

func (dao *UserDAO) validationName(username string) error {

}

func (dao *UserDAO) validationEmail(email string) error {

}

/*ValidationPassword - */
func (dao *UserDAO) validationPassword(password string) error {

}

// /*ChangeModelonModel - изменить модель пользователя на модель пользователя*/
// func (dao *UserDAO) ChangeModelonModel(lastModel, newModel *models.User) (models.User){
// 	var updatedModel models.User
// 	updatedModel.ID = lastModel.ID,
// 	updatedModel.Created = lastModel.Created

// }

// func (dao *UserDAO) changeUserField(modelFieldtype int, modelUser *models.User, lastfield, newfield interface{}){
// 	switch(modelFieldtype){
// 	case 0: // username
// 		if newfield != "<nil>" &&  newfield != "" && len(newfield) >= 5 {
// 			modelUser.Name = newfield
// 		} else {
// 			modelUser.Name = lastfield
// 		}
// 		return

// 	case 1: // email
// 		// TODO: add validation for email
// 		if newfield != "<nil>" &&  newfield != "" && len(newfield) >= 5 {
// 			modelUser.Name = newfield
// 		} else {
// 			modelUser.Name = lastfield
// 		}
// 		return

// 	case 2: // password
// 		if newfield != "<nil>" &&  newfield != "" && len(newfield) >= 5 {
// 			modelUser.Name = newfield
// 		} else {
// 			modelUser.Name = lastfield
// 		}
// 		return
// 	}
// }
