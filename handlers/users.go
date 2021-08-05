package handlers

import (
	"log"
	"net/http"
)

type Users struct {
	l *log.Logger
}

func NewUser(l *log.Logger) *Users {
	return &Users{l}
}

func (u *Users) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("All Users")
}
