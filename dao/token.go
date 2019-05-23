package dao

import (
	"errors"
	"nucleous/models"
)

type TokenDAO struct {
	Database *Database
}

const (
	/*TokensCollection - коллекция с токенами*/
	TokensCollection = "tokens"
)

/*New - инициализация новой бд*/
func (dao *TokenDAO) New(database *Database) *TokenDAO {
	dao.Database = database
	return dao
}

/*CreateNewToken - добавление нового токена в бд*/
func (dao *TokenDAO) CreateNewToken(tok *models.Token) (err error) {
	// tok.Created = time.Now()
	// tok.ID = bson.NewObjectId()

	// err = dao.Database.C(TokensCollection).Insert(tok)
	return
}

/*CheckToken - проверка токена в бд*/
func (dao *TokenDAO) CheckToken(tok *models.Token) (token *models.Token, err error) {
	// var model models.Token

	// err = dao.Database.C(TokensCollection).Find({
	// })

	return nil, errors.New("Not implemented")
}

func (dao *TokenDAO) RemoveToken(tok *models.Token) (err error) {
	return errors.New("Not implemented")
}
