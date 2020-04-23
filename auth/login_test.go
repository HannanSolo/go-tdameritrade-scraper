package auth_test

import (
	"testing"

	"github.com/HannanSolo/go-td-ameritrade-api/auth"
)

func TestURLBuilder(t *testing.T) {

	c := auth.Login{ClientID: "fooID", RedirectURI: "http://localhost"}

	expected := "https://auth.tdameritrade.com/oauth?client_id=fooID%40AMER.OAUTHAP&redirect_uri=http%3A%2F%2Flocalhost&response_type=code"

	if c.URL().String() != expected {
		t.Logf("Got :%#v ", c.URL())
		t.Logf("Exp :%#v ", expected)

		t.Error("Client should of contructed URL correctly")
	}
}
