package providers

import "github.com/go-pg/pg"

func NewDbConnection(config *Config) (*pg.DB, error) {
	connection := pg.Connect(&pg.Options{
		User:     config.DBUser,
		Password: config.DBPass,
		Database: config.DBName,
	})

	if _, err := connection.ExecOne("SELECT 1"); err != nil {
		return nil, err
	}

	return connection, nil
}
