package user

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"labix.org/v2/mgo"
)

var (
	ssName = "current-user"
	store  = sessions.NewCookieStore([]byte("secret-session"))
)

type LocalUser struct {
	Name     string `bson:"name"`
	UserName string `bson:"username"`
	Email    string `bson:"email"`
	Role     string `bson:"role"`
	Password string `bson:"password"`
	Salt     string `bson:"salt"`
}

type User interface {
	IsAuthenticated() bool

	Logout()

	GetById(id interface{}) error

	ProfileURL() string

	Valid(code interface{}) bool
}

// func (u *User) FindUserByUsername(username string) *User {
// 	// TODO
// }

// func (u *User) FindUserBy() {

// }

// func FindUserByUserName(username string) *User {

// }

func Logout() {

	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, ssName)

		delete(session.Values, UserName)
	}
}
