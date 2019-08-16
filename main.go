package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type CustomMux struct {
	http.ServeMux
}

func (c CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if conf.Configuration().Log.Verbose {
		log.Println("Incoming request from", r.Host, "Accessing", r.URL.String())
	}
	c.ServeMux.ServeHTTP(w, r)
}

func main() {
	router := new(CustomMux)
	router.HandleFunc("/", func(w ResponseWriter, r *http.Request) {
		w.Writer([]byte("Hai World"))
	})

	router.HandleFunc("/node", func(w ResponseWriter, r *http.Request) {
		w.Writer([]byte("LOL"))
	})

	server := new(http.Server)
	server.Handler = router
	server.ReadTimeout = conf.Configuration().Server.ReadTimeout * time.Second
	server.WriteTimeout = conf.Configuration().Server.WriteTimeout * time.Second
	server.Addr = fmt.Sprintf(":%d", conf.Configuration().Server.Port)

	if conf.Configuration().Log.Verbose {
		Log.Printf("This server starting at %s \n", serve.Addr)
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
