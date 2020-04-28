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
//   "refresh_token": "U2OH1Db75O4ZicjWs2pmgAlGnYfZUJYXk0hb+F5QAP2PocOSKRPsVUbDsyLsHJ7gLW71m2clCe9xpPvJSfLJwgMGBH+mh9fciWgw/LwDipb41UbvSxMArNZhpf6bXSMdo1QMt6VQiNgKkQc3xOglEkpYSbf+dpAxwnYPuPLxBpjGqqf7wMqB420mTm1cGXphq+MSyzqNomBnMeHlp9hUGUnMrGj3GKliTLl+ewLji6n1vKanDnYtA4c0WMB6CjnvrQoLwznfowT/1gdkmlrnurRxYLAdANOT6nds5+hEZCBGA1KX33dB65X0okOijTdjcWPVc9cLzND9zHjyvHBXNUoSAtSO5Hh7Vu1TopHTbUpGozDRra6TBThP/YgsHVVtMzlpyKODno2GAgk/eoUo7xNB5y3+xWy0a0xzXgfuIvZ7+KI100MQuG4LYrgoVi/JHHvlc2e/jPFhysbWMX3uupZwJTJqeD/qujA1soGdW261VYFeKnVbgSJPmrGPg27BbS9OjK9tQ0AxJueyWl7aRY1ewwZTdemD7QQOfMWucJb9qVIotu0/xzul9pdF+8ZM2l99msoEncyuenEmNtNYgyThmZEEcSwwBT8G4xVWEU/2/wcMUFkE6GqiOitrj368oUnakNnEyl58G/aDv4Im5ka3mc0eefAKzL0oOTetDLwfGM3iQsOVPYUg7amx8xIXrW1CfOYq4lgx67tC+ZbfyGqoITDdRzhWQMZV8yjZQLEhgkSDGPDCmrxpxJM6g+QgIeoYH7MQ/ksitJXDzckaFDUbLqSHpvHk5R1W0DwW6dZ406yO41xH/BPES9dtnuJB+o1hBdZLBMOGPtGhtSvVWy+x9KoYoM4Iq0B4NrqW49guSg/38WZsx0SK0mRmPFg=212FD3x19z9sWBHDJACbC00B75E",
