// +build mem

package mem

import (
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
