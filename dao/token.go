package dao

import (
	"errors"
	"log"
	"nucleous/models"
	"time"

	driver "github.com/arangodb/go-driver"
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

/*GetTokenByUserID - получение токена по идентификатору пользователя*/
func (dao *TokenDAO) GetTokenByUserID(userid string) (*models.Token, error) {

	queriedTokens, err := dao.queryTokenByUserID(userid)

	if err != nil {
		return nil, err
	}

	// log.Println(daoName+"Queired tokens: ", queriedTokens)

	if len(queriedTokens) == 1 {
		return &queriedTokens[0], nil
	} else {
		return nil, errors.New("no have tokens in db collection")
	}

}

func (dao *TokenDAO) queryTokenByUserID(userid string) ([]models.Token, error) {
	query := "FOR d IN tokens FILTER d.user_id == @userid RETURN d"
	bindParams := map[string]interface{}{
		"userid": userid,
	}

	cursor, err := (*dao.Database.Database).Query(nil, query, bindParams)
	if err != nil {
		return nil, err
	}

	var documetns []models.Token

	for {
		var document models.Token

		meta, err2 := cursor.ReadDocument(nil, &document)
		if driver.IsNoMoreDocuments(err2) {
			break
		} else if err2 != nil {
			return nil, err2
		}
		if dao.checkValidToken(&document) {
			documetns = append(documetns, document)
		} else {
			dao.RemoveTokenByID(meta.Key)
		}
	}

	return documetns, nil
}

/*checkValidToken - проверка срока жизни токена относительно текущего момента времени*/
func (dao *TokenDAO) checkValidToken(token *models.Token) bool {
	timeExpired := token.Created.Add(token.ExpireTo)
	if timeExpired.After(time.Now()) {
		return true
	}
	return false
}

/*CreateCollection - создание коллекции с токенами*/
func (dao *TokenDAO) CreateCollection() (driver.Collection, error) {
	return (*dao.Database.Database).CreateCollection(nil, tokensCollection, nil)
}

func (dao *TokenDAO) collectionGet() (*driver.Collection, error) {
	coll, err := (*dao.Database.Database).Collection(nil, tokensCollection)
	if err != nil {
		coll, err := dao.CreateCollection()
		return &coll, err
	}

	return &coll, nil
}

/*CreateNewToken - добавление нового токена в бд*/
func (dao *TokenDAO) CreateNewToken(tok *models.Token) (*models.Token, error) {

	col, err := (*dao.Database.Database).Collection(nil, tokensCollection)

	if err != nil {
		coll, err2 := dao.CreateCollection()
		if err2 != nil {
			return nil, err2
		}

		meta, err2 := coll.CreateDocument(nil, tok)
		if err2 != nil {
			return nil, err2
		}

		log.Println(daoName+" meta information: ", meta)
		tok.ID = meta.Key

		return tok, nil
	}

	meta, err2 := col.CreateDocument(nil, tok)
	if err2 != nil {
		return nil, err2
	}

	log.Println(daoName+" meta information: ", meta)
	tok.ID = meta.Key
	return tok, nil
}

func (dao *TokenDAO) queiredTokenByValue(tokenValue string) ([]models.Token, error) {
	query := "FOR d IN tokens FILTER d.token_value == @token RETURN d"
	bindParams := map[string]interface{}{
		"token": tokenValue,
	}

	cursor, err := (*dao.Database.Database).Query(nil, query, bindParams)
	if err != nil {
		return nil, err
	}

	var documetns []models.Token

	for {
		var document models.Token

		meta, err2 := cursor.ReadDocument(nil, &document)
		if driver.IsNoMoreDocuments(err2) {
			break
		} else if err2 != nil {
			return nil, err2
		}
		if dao.checkValidToken(&document) {
			documetns = append(documetns, document)
		} else {
			dao.RemoveTokenByID(meta.Key)
		}
	}

	return documetns, nil
}

/*CheckToken - проверка токена в бд*/
func (dao *TokenDAO) CheckToken(tokValue string) (token *models.Token, err error) {
	// var model models.Token

	// err = dao.Database.C(TokensCollection).Find({
	// })

	queriedTokens, err := dao.queiredTokenByValue(tokValue)
	if err != nil {
		return nil, err
	}

	if len(queriedTokens) == 1 {
		return &queriedTokens[0], nil
	} else {
		return nil, errors.New("token not exist")
	}

}

/*RemoveTokenByID - удаленние токена по его идентификатору в бд*/
func (dao *TokenDAO) RemoveTokenByID(tokenID string) error {
	coll, err := dao.collectionGet()
	if err != nil {
		return err
	}

	_, err2 := (*coll).RemoveDocument(nil, tokenID)
	if err2 != nil {
		return err2
	}
	return nil
}

/*RemoveToken - удаление токена по его идентификатору*/
func (dao *TokenDAO) RemoveToken(tok *models.Token) (err error) {
	return errors.New("Not implemented")
}
