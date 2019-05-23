package dao

import (
	"errors"
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
)

/*ConnectionUser - подключение пользователя к бд*/
func (db *Database) ConnectionUser(config *configuration.DatabaseConfiguration) error {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{config.Address},
	})

	if err != nil {
		return err
	}

	connectedUser, err2 := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(config.UserName, config.Password),
	})

	if err2 != nil {
		return err2
	}

	db.ConnectUser = &connectedUser

	return nil
}

/*ConnectToBD - подключение к бд*/
func (db *Database) ConnectToBD() error {
	if db.ConnectUser == nil {
		return errors.New("user can not connect to db")
	}

	database, err2 := (*db.ConnectUser).Database(nil, DatabaseName)
	if err2 != nil {
		return err2
	}

	db.Database = &database
	return nil
}
