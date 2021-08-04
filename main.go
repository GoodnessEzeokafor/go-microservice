package main

import (
	"log"
	"github.com/GoodnessEzeokafor/go-microservice/handlers"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", sm)
	http.ListenAndServe(":9090", nil)
}

// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
// 	log.Println("Working on Go Microservice")
// 	d, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		rw.WriteHeader(http.StatusBadRequest)
// 		rw.Write([]byte("Oops"))
// 		return
// 	}
// 	log.Printf("Data '%s'\n", d)
// })
// http.HandleFunc("/user", func(rw http.ResponseWriter, r *http.Request) {
// 	log.Println("user route")
// })
// http.HandleFunc("/logs", func(rw http.ResponseWriter, r *http.Request) {
// 	log.Println("log route")
// })
