package judge

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/superboum/moolinet/lib/sandbox"
	"github.com/superboum/moolinet/lib/tasks"
)

// Challenge holds the data related to a particular challenge.
type Challenge struct {
	Slug        string
	Title       string
	Description string
	Body        string
	Docker      sandbox.DockerSandboxConfig `json:",omitempty"`
	Template    tasks.JobTemplate           `json:",omitempty"`
}

// NewChallengeFromJSON unmarshals a Challenge from a JSON input stream.
func NewChallengeFromJSON(reader io.Reader) (*Challenge, error) {
	decoder := json.NewDecoder(reader)

	chal := new(Challenge)
	err := decoder.Decode(&chal)
	if err != nil {
		return nil, err
	}

	return chal, nil
}

// GetPublicChallenge returns a copy of a Challenge with secret information removed.
func (c *Challenge) GetPublicChallenge() *Challenge {
	newChal := new(Challenge)
	newChal.Slug = c.Slug
	newChal.Title = c.Title
	newChal.Description = c.Description
	newChal.Body = c.Body

	return newChal
}

// LoadChallengesFromPath returns a map of challenge, a potential blocking error and a list of warning caused by unreadable challenges
func LoadChallengesFromPath(challengePath string) (res map[string]*Challenge, loadErrors []error, err error) {
	res = make(map[string]*Challenge)

	files, err := ioutil.ReadDir(challengePath)
	if err != nil {
		return nil, nil, err
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

	return res, loadErrors, nil
}

// GeneratePublicChallenges returns a slice of public Challenges from a map of private Challenges.
func GeneratePublicChallenges(challenges map[string]*Challenge) []*Challenge {
	res := make([]*Challenge, 0, len(challenges))

	for _, val := range challenges {
		res = append(res, val.GetPublicChallenge())
	}

	return res
}
