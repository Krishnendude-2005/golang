package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

const createTableTwo = `CREATE TABLE IF NOT EXISTS users
(
    UserID          int         NOT NULL PRIMARY KEY,
    TaskName varchar(255)
);`

func create_users_table() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			// write your migrations here
			_, err := d.SQL.Exec(createTableTwo)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
