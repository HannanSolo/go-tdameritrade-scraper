package auth_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/HannanSolo/go-td-ameritrade-api/auth"
)

func TestTokenBuilder(t *testing.T) {

	body := []byte(`{
		"access_token":             "seagfs4tg4eewegaa2eaw3rt34g",
		"refresh_token":            "U2OH1wdfwafafawhewr32pmgA",
		"scope":                    "PlaceTrades AccountAccess MoveMoney",
		"expires_in":               1800,
		"refresh_token_expires_in": 7776000,
		"token_type":               "Bearer"}`)

	var response auth.TokenResponse

	json.Unmarshal(body, &response)

	if response.AccessToken != "seagfs4tg4eewegaa2eaw3rt34g" {
		t.Errorf("AccessToken should have been deserialized from access_token property")
	}
	if response.RefreshToken != "U2OH1wdfwafafawhewr32pmgA" {
		t.Errorf("RefreshToken should have been deserialized from refresh_token property")
	}

	t0 := time.Now()
	access, refresh, err := response.Tokens()

	if err != nil {
		t.Errorf("Token extraction should not have returned a error %v", err)
	}

	if access.Code != response.AccessToken {
		t.Errorf("AccessToken should have been assigned right code.")
	}
	if refresh.Code != response.RefreshToken {
		t.Errorf("RefreshToken should have been assigned right code.")
	}
	//time tests

	if access.Expired() {
		t.Error("Token should still be valid")
	}
	if refresh.Expired() {
		t.Error("Token should still be valid")
	}

	if access.Expiration.Sub(t0)-time.Minute*30 > time.Second {
		t.Error("AccessToken should expire 30 minutes from now")
	}
	if refresh.Expiration.Sub(t0)-time.Hour*24*90 > time.Second {
		t.Error("RefreshToken should expire 90 days from now")
	}

}

func TestError(t *testing.T) {

	body := []byte(`{
		"access_token":             "",
		"refresh_token":            "",
		"scope":                    "PlaceTrades AccountAccess MoveMoney",
		"expires_in":               1800,
		"refresh_token_expires_in": 7776000,
		"token_type":               "Bearer"}`)

	var response auth.TokenResponse

	json.Unmarshal(body, &response)

	if _, _, err := response.Tokens(); err == nil {
		t.Error("Tokens should have returned an error")
	}

}

// take above example response and test it
// Deserialize json object using the json library and put into token response struct struct
// Ensure the tokens method converts the time correctly and test for edge cases. (ex, expires in 5 seconds)
// Ensure malform data works
// test worst case
//test extra fields
// add package level function to take an io.Reader contruct a proper json decoder and Unmarshal it to an object
// func NewResp(r io.Reader) (TokenResponse, err)
