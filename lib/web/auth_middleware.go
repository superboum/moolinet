package web

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/superboum/moolinet/lib/persistence"
	"github.com/superboum/moolinet/lib/tools"
)

// ErrNotAuthenticated is returned when we try to get the logged user but no one is authenticated
var ErrNotAuthenticated = errors.New("Not authenticated")

// ErrWrongUserType is returned when the thing stored in the session is not a User
var ErrWrongUserType = errors.New("Wrong User Type")

// AuthMiddleware is a middleware to check authentication
type AuthMiddleware struct {
	Store *sessions.CookieStore
}

// NewAuthMiddleware returns a new auth middleware
func NewAuthMiddleware() *AuthMiddleware {
	a := &AuthMiddleware{
		Store: sessions.NewCookieStore([]byte(tools.GeneralConfig.CookieSecretKey)),
	}

	gob.Register(persistence.User{})
	return a
}

// GetUser return the current user
func (a *AuthMiddleware) GetUser(req *http.Request) (*persistence.User, error) {
	session, err := a.Store.Get(req, "auth")
	if err != nil {
		return nil, err
	}

	val, ok := session.Values["user"]
	if !ok {
		return nil, ErrNotAuthenticated
	}

	log.Println(val)
	user, ok := val.(persistence.User)
	if !ok {
		return nil, ErrWrongUserType
	}

	return &user, nil
}

// SetUser define the logged user stored in the session
func (a *AuthMiddleware) SetUser(res http.ResponseWriter, req *http.Request, u *persistence.User) error {
	session, err := a.Store.Get(req, "auth")
	if err != nil {
		return err
	}

	session.Values["user"] = u
	err = session.Save(req, res)

	return err
}

// CheckAuthentication checks that an user is authenticated (with a session)
func (a *AuthMiddleware) CheckAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, err := a.GetUser(req)
		if err == ErrNotAuthenticated {
			encoder := json.NewEncoder(res)
			res.WriteHeader(401)
			checkEncode(encoder.Encode(APIError{"Wrong credentials", "You should log on the API before"}))
		} else if err != nil {
			log.Println(err)
			encoder := json.NewEncoder(res)
			res.WriteHeader(500)
			checkEncode(encoder.Encode(APIError{"Internal error", "Contact an administrator"}))
		} else {
			next.ServeHTTP(res, req)
		}
	})
}
