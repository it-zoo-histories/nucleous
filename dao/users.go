package dao

import (
	"errors"
	"log"
	"nucleous/models"
	"nucleous/payloads"
	"time"
)

/*UserDAO - dao для работы с коллекцией пользователей*/
type UserDAO struct {
	Database *Database
}

const (
	usersCollection = "users"
	dao2Name        = "[NUCLEOUS:USERS]: "
)

/*New - инициализация userdao инстанса*/
func (dao *UserDAO) New(database *Database) *UserDAO {
	dao.Database = database
	return dao
}

/*CheckExistCollection - проверка наличия коллекции в бд*/
func (dao *UserDAO) CheckExistCollection() error {
	_, err := (*dao.Database.Database).CollectionExists(nil, usersCollection)
	return err
}

/*CreateCollection - создание коллекции с токенами*/
func (dao *UserDAO) CreateCollection() error {
	_, err := (*dao.Database.Database).CreateCollection(nil, usersCollection, nil)
	return err
}

/*CreateNewUser - создание нового пользователя*/
func (dao *UserDAO) CreateNewUser(user *models.User) (err error) {
	user.Created = time.Now()
	user.Verificated = false

	coll, err := (*dao.Database.Database).Collection(nil, usersCollection)
	if err != nil {
		return err
	}

	meta, err2 := coll.CreateDocument(nil, user)
	if err2 != nil {
		return err2
	}

	log.Println(dao2Name+" meta infor after creating: ", meta)

	// err = dao.Database.C(UsersCollection).Insert(&user)
	return
}

/*UpdateUser - обновление данных пользователя по его идентификатору*/
func (dao *UserDAO) UpdateUser(userid string, newpayload *payloads.UserUpdatePayload) (updated *models.User, err error) {
	// iduser := bson.ObjectIdHex(userid)

	// var modelUser models.User

	// err = dao.Database.C(UsersCollection).Find(bson.M{
	// "_id": iduser,
	// }).One(&modelUser)

	// if err != nil {
	// return &models.User{}, err
	// }

	// if errCheckUserVerification := dao.checkVerificationUser(modelUser); errCheckUserVerification != nil {
	// return nil, errCheckUserVerification
	// }
	return nil, nil
}

/*setupNewPayload - установка новых значений пользователю*/
func (dao *UserDAO) setupNewPayload(userInDb *models.User, payload *payloads.UserUpdatePayload) error {
	return nil
}

/*checkVerificationUser - проверка подтверждённия аккаунта пользователя*/
func (dao *UserDAO) checkVerificationUser(userModel models.User) error {
	if userModel.Verificated {
		return nil
	} else {
		return errors.New("user can not change user information")
	}
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
