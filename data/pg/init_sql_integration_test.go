// +build integration
// +build !mgo

package pg

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/trimcao/go-saas/data/model"
)

var db *sql.DB

func TestMain(m *testing.M) {
	conn, err := model.Open("postgres",
		"postgres://postgres:dbpwd@localhost/gosaas?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	db = conn

	os.Exit(m.Run())
}
