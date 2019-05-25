package dao

import (
	"errors"
	"log"
	"nucleous/models"
	"nucleous/payloads"
	"time"

	driver "github.com/arangodb/go-driver"
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
func (dao *UserDAO) CreateCollection() (driver.Collection, error) {
	return (*dao.Database.Database).CreateCollection(nil, usersCollection, nil)
}

func (dao *UserDAO) collectionGet() (*driver.Collection, error) {
	coll, err := (*dao.Database.Database).Collection(nil, usersCollection)
	if err != nil {
		coll, err := dao.CreateCollection()
		return &coll, err
	}

	return &coll, nil
}

/*CreateNewUser - создание нового пользователя*/
func (dao *UserDAO) CreateNewUser(userPay payloads.CreateUserPayload) (err error) {
	var user models.User
	user.Email = userPay.Email
	user.Created = time.Now()
	user.Verificated = false
	user.Name = userPay.Username
	user.Avatar = userPay.Avatar
	user.HashAndSalt(userPay.Password)

	coll, err := dao.collectionGet()
	if err != nil {
		return err
	}

	meta, err2 := (*coll).CreateDocument(nil, user)
	if err2 != nil {
		return err2
	}

	log.Println(dao2Name+" meta infor after creating: ", meta)

	// err = dao.Database.C(UsersCollection).Insert(&user)
	return err2
}

/*FindUserByUserID - поиск пользователя по его идентификатору*/
func (dao *UserDAO) FindUserByUserID(userid string) (*models.User, error) {
	var document models.User

	coll, err := dao.collectionGet()
	if err != nil {
		return nil, err
	}

	meta, err2 := (*coll).ReadDocument(nil, userid, &document)

	log.Println(dao2Name+"meta info after getting document: ", meta)

	return &document, err2
}

/*UpdateUserVerification - обновление верификации пользователя*/
func (dao *UserDAO) UpdateUserVerification(userid string, update bool) error {
	var document models.User
	coll, err := dao.collectionGet()
	if err != nil {
		return err
	}

	_, err2 := (*coll).ReadDocument(nil, userid, &document)
	if err2 != nil {
		return err2
	}

	document.Verificated = update

	_, err3 := (*coll).UpdateDocument(nil, userid, document)

	return err3
}

/*UpdateUser - обновление данных пользователя по его идентификатору*/
func (dao *UserDAO) UpdateUser(userid string, newpayload *payloads.UserUpdatePayload) (updated *models.User, err error) {
	// iduser := bson.ObjectIdHex(userid)

	var document models.User
	coll, err := dao.collectionGet()
	if err != nil {
		return nil, err
	}

	_, err2 := (*coll).ReadDocument(nil, userid, &document)
	if err2 != nil {
		return nil, err2
	}

	if newpayload.ValidateUsername {
		document.Name = newpayload.UserName
	}
	if newpayload.ValidatePassword {
		document.HashAndSalt(newpayload.Password)
	}
	if newpayload.ValidateEmail {
		document.Email = newpayload.Email
		document.Verificated = false
	}
	if newpayload.ValidateAvatar {
		document.Avatar = newpayload.Avatar
	}
	if newpayload.ValidateVerification {
		document.Verificated = newpayload.Verification
	}

	_, err3 := (*coll).UpdateDocument(nil, userid, document)

	if err3 != nil {
		return nil, err
	}

	return &document, nil
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
