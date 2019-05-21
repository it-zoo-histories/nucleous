package dao

import (
	"errors"
	"nucleous/models"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TokenDAO struct {
	Database *mgo.Database
}

const (
	/*TokensCollection - коллекция с токенами*/
	TokensCollection = "tokens"
)

/*CreateNewToken - добавление нового токена в бд*/
func (dao *TokenDAO) CreateNewToken(tok *models.Token) (err error) {
	tok.Created = time.Now()
	tok.ID = bson.NewObjectId()

	err = dao.Database.C(TokensCollection).Insert(tok)
	return
}

/*CheckToken - проверка токена в бд*/
func (dao *TokenDAO) CheckToken(tok *models.Token) (token *models.Token, err error) {
	var model models.Token

	// err = dao.Database.C(TokensCollection).Find({
	// })

	return nil, errors.New("Not implemented")
}

func (dao *TokenDAO) RemoveToken(tok *models.Token) (err error) {
	return errors.New("Not implemented")
}
