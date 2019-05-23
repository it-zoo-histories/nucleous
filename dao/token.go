package dao

import (
	"errors"
	"nucleous/models"
)

/*TokenDAO - */
type TokenDAO struct {
	Database *Database
}

const (
	/*TokensCollection - коллекция с токенами*/
	tokensCollection = "tokens"
)

/*New - инициализация новой бд*/
func (dao *TokenDAO) New(database *Database) *TokenDAO {
	dao.Database = database
	return dao
}

/*CheckExistCollection - проверка наличия коллекции в бд*/
func (dao *TokenDAO) CheckExistCollection() error {
	_, err := (*dao.Database.Database).CollectionExists(nil, tokensCollection)
	return err
}

/*CreateCollection - создание коллекции с токенами*/
func (dao *TokenDAO) CreateCollection() error {
	_, err := (*dao.Database.Database).CreateCollection(nil, tokensCollection, nil)
	return err
}

/*CreateNewToken - добавление нового токена в бд*/
func (dao *TokenDAO) CreateNewToken(tok *models.Token) (err error) {
	// tok.Created = time.Now()
	// tok.ID = bson.NewObjectId()
	// context := context.Background()
	// err = dao.Database.C(TokensCollection).Insert(tok)
	return
}

/*CheckToken - проверка токена в бд*/
func (dao *TokenDAO) CheckToken(tokValue string) (token *models.Token, err error) {
	// var model models.Token

	// err = dao.Database.C(TokensCollection).Find({
	// })

	return nil, errors.New("Not implemented")
}

/*RemoveToken - удаление токена по его идентификатору*/
func (dao *TokenDAO) RemoveToken(tok *models.Token) (err error) {
	return errors.New("Not implemented")
}
