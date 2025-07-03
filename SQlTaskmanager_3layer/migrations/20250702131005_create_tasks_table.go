package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

const createTable = `CREATE TABLE IF NOT EXISTS tasks
(
    ID          int         NOT NULL PRIMARY KEY,
    Description varchar(255),
    Status      boolean,
    UserID      int
);`

func create_tasks_table() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			// write your migrations here
			_, err := d.SQL.Exec(createTable)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
