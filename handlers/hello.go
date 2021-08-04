package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}
func NewHello(l *log.Logger) *Hello{
	return &Hello{l}
}
// define a method on a struct

func (h *Hello) serveHttp(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Working on Go Microservice")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Oops"))
		return
	}
	log.Printf("Data '%s'\n", d)
}
