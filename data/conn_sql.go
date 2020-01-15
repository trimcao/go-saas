// +build !mgo,!mem

package data

import (
	"github.com/trimcao/go-saas/data/model"
	"github.com/trimcao/go-saas/data/pg"
)

func (db *DB) Open(driverName, dataSourceName string) error {
	conn, err := model.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}

	db.Users = &pg.Users{DB: conn}

	// initialize services
	db.Connection = conn
	return nil
}
