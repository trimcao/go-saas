// +build !mgo

package data

import (
	_ "github.com/lib/pq"
	"github.com/trimcao/go-saas/data/model"
)

func (db *DB) Open(driverName, dataSourceName string) error {
	conn, err := model.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}

	// initialize services
	db.Connection = conn
	return nil
}
