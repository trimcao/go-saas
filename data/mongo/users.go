package mongo

import (
	"github.com/trimcao/go-saas/data/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Users struct {
	DB *mgo.Database
}

func (u *Users) GetDetail(id model.Key) (*model.User, error) {
	var user model.User
	where := bson.M{"_id": id}
	if err := u.DB.C("users").Find(where).One(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *Users) RefreshSession(s *mgo.Session, dbName string) {
	u.DB = s.Copy().DB(dbName)
}

func (u *Users) SignUp(email, password string) (*model.Account, error) {
	acct := model.Account{ID: bson.NewObjectId(), Email: email}
	acct.Users = append(acct.Users, model.User{ID: bson.NewObjectId(), Email: email, Password: password})

	if err := u.DB.C("users").Insert(acct); err != nil {
		return nil, err
	}
	return &acct, nil
}
