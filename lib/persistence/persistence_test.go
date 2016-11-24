package persistence

import (
	"os"
	"testing"

	"github.com/superboum/moolinet/lib/tools"
)

func TestUser(t *testing.T) {
	_ = os.Remove("moolinet.db")
	err := tools.LoadGeneralConfigFromFile("../../moolinet.json")
	if err != nil {
		t.Error(err)
	}

	err = InitDatabase()
	if err != nil {
		t.Error(err)
	}

	u, err := NewUser("test", "coucou")
	if err != nil {
		t.Error(err)
	}
	if u.Username != "test" {
		t.Error("Wrong user name")
	}

	u, err = LoginUser("test", "coucou")
	if err != nil {
		t.Error(err)
	}
	if u.Username != "test" {
		t.Error("Wrong user name")
	}

	u, err = LoginUser("test2", "coucou")
	if err.Error() != "Wrong credentials" {
		t.Error("Expecting 'Wrong credentials' but get", err.Error())
	}

	u, err = LoginUser("test", "coucou2")
	if err.Error() != "Wrong credentials" {
		t.Error("Expecting 'Wrong credentials' but get", err.Error())
	}
}
