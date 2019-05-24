package dao

import (
	"errors"
	"log"
	"nucleous/models"
	"time"
)

/*TokenDAO - */
type TokenDAO struct {
	Database *Database
}

const (
	/*TokensCollection - коллекция с токенами*/
	tokensCollection = "tokens"
	daoName          = "[NUCLEOUS:TOKENS]: "
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

	tok.Created = time.Now()
	col, err := (*dao.Database.Database).Collection(nil, tokensCollection)

	if err != nil {
		return err
	}

	meta, err2 := col.CreateDocument(nil, tok)
	if err2 != nil {
		return err2
	}

	log.Println(daoName+" meta information: ", meta)
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
