package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker service in port %s", webPort)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes() ,
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}