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

	_, err = NewUser("test", "coucou")
	if err != nil {
		t.Error(err)
	}
}
