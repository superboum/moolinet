package persistence

import (
	"time"

	// Used for the driver database/sql
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// User struct containing user informations
type User struct {
	Username string
	password string
	Created  time.Time
}

// NewUser creates and adds a new user to the database
func NewUser(username string, password string) (*User, error) {
	u := &User{Username: username}
	err := u.SetPassword(password)
	if err != nil {
		return nil, err
	}

	stmt, err := DB.Prepare("INSERT INTO user(username, password, created) values(?,?,?)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(u.Username, u.password, "2012-12-09")
	if err != nil {
		return nil, err
	}

	return u, err
}

// LoginUser search a user in the database and return this user if the user exists
func LoginUser(username string, password string) (*User, error) {
	return nil, nil
}

// SetPassword hash users password
func (u *User) SetPassword(password string) error {
	encoded, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err == nil {
		u.password = string(encoded[:])
	}
	return err
}
