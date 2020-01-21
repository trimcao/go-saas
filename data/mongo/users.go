// +build !mem

package mongo

import (
	"fmt"

	"github.com/trimcao/go-saas/data/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Users struct {
	DB *mgo.Database
}

func (u *Users) GetDetail(id model.Key) (*model.Account, error) {
	var acct model.Account
	where := bson.M{"_id": id}
	if err := u.DB.C("users").Find(where).One(&acct); err != nil {
		return nil, err
	}
	return &acct, nil
}

func (u *Users) RefreshSession(s *mgo.Session, dbName string) {
	u.DB = s.Copy().DB(dbName)
}

func (u *Users) SignUp(email, password string) (*model.Account, error) {
	accountID := bson.NewObjectId()
	acct := model.Account{ID: accountID, Email: email}
	acct.Users = append(acct.Users, model.User{
		ID:       accountID,
		Email:    email,
		Password: password,
		Token:    model.NewToken(accountID),
	})

	if err := u.DB.C("users").Insert(acct); err != nil {
		return nil, err
	}
	return &acct, nil
}

func (u *Users) Auth(accountID, token string, pat bool) (*model.Account, *model.User, error) {
	if bson.IsObjectIdHex(accountID) == false {
		return nil, nil, fmt.Errorf("this account id is invalid %s", accountID)
	}

	id := bson.ObjectIdHex(accountID)

	acct, err := u.GetDetail(id)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot get the account with id %s", id)
	}

	var user model.User
	for _, usr := range acct.Users {
		if pat {
			if usr.Token == token {
				user = usr
				break
			}
		} else {
			for _, at := range usr.AccessTokens {
				if at.Token == token {
					user = usr
					break
				}
			}
		}
	}
	return acct, &user, nil
}
