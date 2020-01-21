package data

import "github.com/trimcao/go-saas/data/model"

type SessionRefresher interface {
	RefreshSession(*model.Connection, string)
}

type UserServices interface {
	SessionRefresher
	GetDetail(id model.Key) (*model.Account, error)
	SignUp(email, password string) (*model.Account, error)
	AddToken(accountID, userID model.Key, name string) (*model.AccessToken, error)
	RemoveToken(accountID, userID, tokenID model.Key) error
	Auth(accountID, token string, pat bool) (*model.Account, *model.User, error)
}

type DB struct {
	DatabaseName string
	Connection   *model.Connection
	CopySession  bool

	Users UserServices
}
