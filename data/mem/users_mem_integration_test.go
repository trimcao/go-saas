// +build mem

package mem

import "testing"

func Test_DB_Users_SignUp(t *testing.T) {
	users := Users{}
	// users.RefreshSession(db, dbName)

	acct, err := users.SignUp("unit@test.com", "unittest")
	if err != nil {
		t.Error("unable to create new account", err)
	} else if acct.Email != "unit@test.com" {
		t.Error("account email is not correct expected unit@test.com got", acct.Email)
	} else if len(acct.Users) != 1 {
		t.Error("account has no users expected 1 got", len(acct.Users))
	} else if acct.Users[0].Email != "unit@test.com" {
		t.Error("user email is not correct expected unit@test.com got", acct.Users[0].Email)
	}
}
