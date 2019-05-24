package dao

import (
	"context"
	"errors"
	"log"
	"nucleous/configuration"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

/*Database - обёртка над подключённым юзером к бд*/
type Database struct {
	ConnectUser *driver.Client
	Database    *driver.Database
}

const (
	// DatabaseName -
	DatabaseName = "Client"
	nameServer   = "[NUCLEOUS]: "
)

/*ConnectionUser - подключение пользователя к бд*/
func (db *Database) ConnectionUser(config *configuration.DatabaseConfiguration) error {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{config.Address},
	})

	if err != nil {
		log.Println(nameServer+"Error code: ", err.Error())
		return err
	}

	log.Println(nameServer + "success connect to database")

	connectedUser, err2 := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(config.UserName, config.Password),
	})

	if err2 != nil {
		log.Println(nameServer+"Error code: ", err2.Error())
		return err2
	}

	log.Println(nameServer + "success connect user to bd")

	db.ConnectUser = &connectedUser

	log.Println(nameServer, *db)

	return nil
}

/*ConnectToBD - подключение к бд*/
func (db *Database) ConnectToBD() error {
	if db.ConnectUser == nil {
		return errors.New("user can not connect to db")
	}

	log.Println(nameServer+"success connect to bd", db)
	ctx := context.Background()

	database, err2 := (*db.ConnectUser).Database(ctx, DatabaseName)
	if err2 != nil {
		log.Println(nameServer+" error code: ", err2.Error())
		return err2
	}

	db.Database = &database
	return nil
}
