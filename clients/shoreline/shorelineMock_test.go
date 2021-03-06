package shoreline

import (
	"testing"
)

func TestMock(t *testing.T) {

	const TOKEN_MOCK = "this is a token"

	client := NewMock(TOKEN_MOCK)

	if err := client.Start(); err != nil {
		t.Errorf("Failed start with error[%v]", err)
	}

	if tok := client.TokenProvide(); tok != TOKEN_MOCK {
		t.Errorf("Unexpected token[%s]", tok)
	}

	if usr, token, err := client.Login("billy", "howdy"); err != nil {
		t.Errorf("Failed start with error[%v]", err)
	} else {
		if usr == nil {
			t.Error("Should give us a fake usr details")
		}
		if token == "" {
			t.Error("Should give us a fake token")
		}
	}

	if checkedTd := client.CheckToken(TOKEN_MOCK); checkedTd == nil {
		t.Error("Should give us token data")
	}

	if usr, _ := client.GetUser("billy@howdy.org", TOKEN_MOCK); usr == nil {
		t.Error("Should give us a mock user")
	}

	user := UserUpdate{UserData: UserData{UserID: "123", UserName: "name", Emails: []string{"an@email.org"}}, Password: "myN3wPw"}

	if err := client.UpdateUser(user, TOKEN_MOCK); err != nil {
		t.Error("Should return no error on success")
	}

	if host := client.getHost(); host == nil {
		t.Error("Should give us a fake host")
	}

	if err := client.serverLogin(); err != nil {
		t.Error("Should not return err")
	}

}
