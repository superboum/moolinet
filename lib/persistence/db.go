package persistence

import (
	"database/sql"

	// That's a SQL driver we'll use with the database/sql API
	_ "github.com/mattn/go-sqlite3"
	"github.com/superboum/moolinet/lib/tools"
)

// DB is the database object shared by everyone
var DB *sql.DB

// InitDatabase return a singleton database object
func InitDatabase() error {
	var err error
	DB, err = sql.Open("sqlite3", tools.GeneralConfig.DatabasePath)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
CREATE TABLE IF NOT EXISTS user (
  username VARCHAR(255) PRIMARY KEY,
	password VARCHAR(255) NOT NULL,
	created DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS job (
	uuid VARCHAR(255) PRIMARY KEY,
	challenge VARCHAR(255) NOT NULL,
  username VARCHAR(255) NOT NULL,
	status INTEGER NOT NULL,
	created DATETIME NOT NULL,
	FOREIGN KEY(username) REFERENCES user(username)
);
	`)

	return err
}
