package data

import "github.com/trimcao/go-saas/data/model"

type SessionRefresher interface {
	RefreshSession(*model.Connection, string)
}

type UserServices interface {
	SessionRefresher
	GetDetail(id model.Key) (*model.User, error)
}

type DB struct {
	DatabaseName string
	Connection   *model.Connection
	CopySession  bool

	Users UserServices
}
