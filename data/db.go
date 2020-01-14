package data

import "github.com/trimcao/go-saas/data/model"

type DB struct {
	DatabaseName string
	Connection   *model.Connection
}
