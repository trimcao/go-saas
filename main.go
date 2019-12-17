package main

import (
	"log"
	"net/http"

	"github.com/trimcao/go-saas/controllers"
)

func main() {
	api := &controllers.API{}

	if err := http.ListenAndServe(":8080", api); err != nil {
		log.Println(err)
	}
}
