package web

import (
	"encoding/gob"
	"encoding/json"
	"log"
	"net/http"

	"github.com/superboum/moolinet/lib/persistence"
)

// AuthController is a struct used to managed the controller
type AuthController struct {
	auth    *AuthMiddleware
	baseURL string
}

// NewAuthController returns a new auth controller
func NewAuthController(baseURL string, auth *AuthMiddleware) *AuthController {
	a := &AuthController{
		auth:    auth,
		baseURL: baseURL,
	}

	gob.Register(persistence.User{})
	return a
}

type credentials struct {
	Username string
	Password string
}

// Const for HTTP Methods
const (
	HTTPMethodPost = "POST"
)

// ServeHTTP handles authentication requests
func (a *AuthController) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	action := req.URL.Path[len(a.baseURL):]
	switch {
	case req.Method == HTTPMethodPost && action == "login":
		a.handleLogin(res, req)
		return
	case req.Method == HTTPMethodPost && action == "logout":
		a.handleLogout(res, req)
		return
	case req.Method == HTTPMethodPost && action == "register":
		a.handleRegister(res, req)
		return
	}

	// Handle error
	a.handle404(res, req)
}

func (a *AuthController) getCredentials(res http.ResponseWriter, req *http.Request) (credentials, error) {
	decoder := json.NewDecoder(req.Body)

	creds := credentials{}
	err := decoder.Decode(&creds)

	if err != nil {
		encoder := json.NewEncoder(res)
		res.WriteHeader(400)
		checkEncode(encoder.Encode(APIError{"Your request is malformed", "Put your JSON body in a validator and check the API"}))
	}
	return creds, err
}

func (a *AuthController) setUserToSession(res http.ResponseWriter, req *http.Request, u *persistence.User) error {
	err := a.auth.SetUser(res, req, u)
	if err != nil {
		res.WriteHeader(500)
		encoder := json.NewEncoder(res)
		checkEncode(encoder.Encode(APIError{"Internal error", "Contact an administrator"}))
		log.Println(err)
	}
	return err

}

func (a *AuthController) handleLogin(res http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(res)
	creds, err := a.getCredentials(res, req)
	if err != nil {
		return
	}

	u, err := persistence.LoginUser(creds.Username, creds.Password)
	if err != nil && err.Error() == "Wrong credentials" {
		res.WriteHeader(401)
		checkEncode(encoder.Encode(APIError{"Wrong credentials", "Check your username and your password"}))
		return
	} else if err != nil {
		res.WriteHeader(500)
		checkEncode(encoder.Encode(APIError{"Internal error", "Contact an administrator"}))
		log.Println(err)
		return
	}

	if a.setUserToSession(res, req, u) != nil {
		return
	}

	checkEncode(encoder.Encode(u))
}

func (a *AuthController) handleLogout(res http.ResponseWriter, req *http.Request) {
	if a.setUserToSession(res, req, nil) != nil {
		return
	}
	res.WriteHeader(204)
}

func (a *AuthController) handleRegister(res http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(res)
	creds, err := a.getCredentials(res, req)
	if err != nil {
		return
	}

	u, err := persistence.NewUser(creds.Username, creds.Password)
	if err != nil {
		res.WriteHeader(500)
		checkEncode(encoder.Encode(APIError{"Internal error", "Contact an administrator"}))
		log.Println(err)
		return
	}

	if a.setUserToSession(res, req, u) != nil {
		return
	}

	checkEncode(encoder.Encode(u))
}

func (a *AuthController) handle404(res http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(res)
	res.WriteHeader(404)
	checkEncode(encoder.Encode(APIError{"Action unknown", "Check your request"}))
}
