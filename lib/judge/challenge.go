package judge

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/superboum/moolinet/lib/tasks"
)

type Challenge struct {
	Slug        string
	Title       string
	Description string
	Body        string
	Image       string            `json:",omitempty"`
	Template    tasks.JobTemplate `json:",omitempty"`
}

func NewChallengeFromJSON(reader io.Reader) (*Challenge, error) {
	decoder := json.NewDecoder(reader)

	chal := new(Challenge)
	err := decoder.Decode(&chal)
	if err != nil {
		return nil, err
	}

	return chal, nil
}

func (c *Challenge) GetPublicChallenge() *Challenge {
	newChal := new(Challenge)
	newChal.Slug = c.Slug
	newChal.Title = c.Title
	newChal.Description = c.Description
	newChal.Body = c.Body

	return newChal
}

// LoadChallengesFromPath returns a map of challenge, a potential blocking error and a list of warning caused by unreadable challenges
func LoadChallengesFromPath(challengePath string) (map[string]*Challenge, error, []error) {
	res := make(map[string]*Challenge)
	loadErrors := make([]error, 0)

	files, err := ioutil.ReadDir(challengePath)
	if err != nil {
		return nil, err, nil
	}

	for _, file := range files {
		if !file.IsDir() {
			fullpath := path.Join(challengePath, file.Name())
			reader, err := os.Open(fullpath)
			if err != nil {
				log.Println("Unable to open", fullpath, "error:", err)
				loadErrors = append(loadErrors, err)
				continue
			}

			chal, err := NewChallengeFromJSON(reader)
			if err != nil {
				log.Println("Unable to parse", fullpath, "error:", err)
				loadErrors = append(loadErrors, err)
				continue
			}
			res[chal.Slug] = chal
		}
	}

	return res, nil, loadErrors
}

func GeneratePublicChallenges(challenges map[string]*Challenge) []*Challenge {
	res := make([]*Challenge, 0, len(challenges))

	for _, val := range challenges {
		res = append(res, val.GetPublicChallenge())
	}

	return res
}
