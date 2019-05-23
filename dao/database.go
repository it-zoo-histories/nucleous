package dao

import (
	"nucleous/configuration"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

/*Database - обёртка над подключённым юзером к бд*/
type Database struct {
	ConnectUser *driver.Client
}

/*Connection - подключение к бд*/
func (db *Database) Connection(config *configuration.DatabaseConfiguration) error {
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
