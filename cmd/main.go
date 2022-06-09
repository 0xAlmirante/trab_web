package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{
		Addr:    ":4000",
		Handler: Rotas(),
	}

	erro := srv.ListenAndServe()
	if erro != nil {
		log.Fatal(erro)
	}
}
