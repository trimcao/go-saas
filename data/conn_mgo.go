// +build !mem,!sql

package data

import "github.com/trimcao/go-saas/data/model"

func (db *DB) Open(driverName, dataSourceName string) error {
	conn, err := model.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}

	// for mongo, we need to copy the connection session at each requests
	// this is done in our api's ServeHTTP
	db.CopySession = true

	// initialize services
	db.Connection = conn
	return nil
}
