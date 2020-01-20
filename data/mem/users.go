// +build mem

package mem

import "github.com/trimcao/go-saas/data/model"

type Users struct {
	Store []model.User
}

func (u *Users) GetDetail(id model.Key) (*model.User, error) {
	var user model.User
	for _, usr := range u.Store {
		if usr.ID == id {
			user = usr
			break
		}
	}
	return &user, nil
}

func (u *Users) RefreshSession(conn *bool, dbName string) {
	u.Store = append(u.Store, model.User{ID: 1, Email: "test@domain.com"})
}
