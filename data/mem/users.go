// +build mem

package mem

import (
	"fmt"

	"github.com/trimcao/go-saas/data/model"
)

type Users struct {
	Store []model.Account
}

func (u *Users) GetDetail(id model.Key) (*model.Account, error) {
	var account model.Account
	for _, acct := range u.Store {
		if acct.ID == id {
			account = acct
			break
		}
	}
	return &account, nil
}

func (u *Users) RefreshSession(conn *bool, dbName string) {
	// u.Store = append(u.Store, model.User{ID: 1, Email: "test@domain.com"})
}

func (u *Users) SignUp(email, password string) (*model.Account, error) {
	accountID := int64(len(u.Store))
	acct := model.Account{ID: accountID, Email: email}
	acct.Users = append(acct.Users, model.User{
		ID:       accountID,
		Email:    email,
		Password: password,
		Token:    model.NewToken(accountID),
	})

	u.Store = append(u.Store, acct)

	return &acct, nil
}

func (u *Users) AddToken(accountID, userID model.Key, name string) (*model.AccessToken, error) {
	tok := model.AccessToken{
		ID:    userID * 300,
		Name:  name,
		Token: model.NewToken(accountID),
	}

	for _, acct := range u.Store {
		if acct.ID == accountID {
			for _, usr := range acct.Users {
				if usr.ID == userID {
					usr.AccessTokens = append(usr.AccessTokens, tok)
					return &tok, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("unable to find account %d and user %d", accountID, userID)
}

func (u *Users) RemoveToken(accountID, userID, tokenID model.Key) error {
	for _, acct := range u.Store {
		if acct.ID == accountID {
			for _, usr := range acct.Users {
				if usr.ID == userID {
					usr.AccessTokens = make([]model.AccessToken, 0)
					break
				}
			}
		}
	}
	return nil
}
