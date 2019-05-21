package dao

import (
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
func (dao *TokenDAO) CheckToken(tok *models.Token) (err error) {
	var model models.Token

	err = dao.Database.C(TokensCollection).
}

func (dao *TokenDAO) RemoveToken(tok *models.Token)(err error){
	
}
