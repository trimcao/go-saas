// +build mem

package data

import (
	"github.com/trimcao/go-saas/data/mem"
	"github.com/trimcao/go-saas/data/model"
)

func (db *DB) Open(driverName, dataSourceName string) error {
	_, err := model.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}

	db.Users = &mem.Users{Store: []model.Account{}}

	// initialize services
	// db.Connection = conn
	return nil
}
