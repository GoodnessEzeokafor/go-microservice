package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/GoodnessEzeokafor/go-microservice/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	uh := handlers.NewUser(l)
	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/users", uh)
	sm.Handle("/products", ph)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	l.Println("Recieved terminate, gracefully shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
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
