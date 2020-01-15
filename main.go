package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/trimcao/go-saas/controllers"
	"github.com/trimcao/go-saas/data"
)

func main() {
	dn := flag.String("driver", "postgres", "name of the database driver to use, postgres or mongo are supported")
	ds := flag.String("datasource", "", "database connection string")
	flag.Parse()

	if len(*dn) == 0 || len(*ds) == 0 {
		flag.Usage()
		return
	}

	api := controllers.NewAPI()

	// open the database connection
	db := &data.DB{}
	if err := db.Open(*dn, *ds); err != nil {
		log.Fatal("unable to connect to the database:", err)
	}
	api.DB = db

	log.Println("Server is listening at port 8080")
	if err := http.ListenAndServe(":8080", api); err != nil {
		log.Println(err)
	}
}
