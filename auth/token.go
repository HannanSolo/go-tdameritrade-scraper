package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const tokenEndPoint = "https://api.tdameritrade.com/v1/oauth2/token"

const gracePeriod = 40 * time.Second

// "grant_type":   "authorization_code",
// 		"access_type":  "offline",
// 		"code":         k.AccessToken.Code,
// 		"client_id":    k.ClientID,
// 		"redirect_uri": "http://localhost",

type TokenRequest struct {
	GrantType   string `json:"grant_type"`
	AccessType  string `json:"access_type"`
	Code        string `json:"code"`
	ClientID    string `json:"client_id"`
	RedirectURI string `json:"redirect_uri"`
}

type TokenResponse struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	Scope                 string `json:"scope"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
	TokenType             string `json:"token_type"`
}

type TokenManager struct {
	ClientID     string
	AccessToken  Token
	RefreshToken Token
}

type Token struct {
	Code       string
	Expiration time.Time
}

// Expired checks if a token is expired
func (t Token) Expired() bool {
	return time.Now().Add(gracePeriod).After(t.Expiration)
}

func (t *TokenResponse) Tokens() (access Token, refresh Token, err error) {

	if t.AccessToken == "" || t.RefreshToken == "" {
		err = fmt.Errorf("Missing token in TokenResponse")
		return
	}
	access.Code = t.AccessToken
	refresh.Code = t.RefreshToken

	access.Expiration = time.Now().Add(time.Second * time.Duration(t.ExpiresIn))
	refresh.Expiration = time.Now().Add(time.Second * time.Duration(t.RefreshTokenExpiresIn))

	return
}

func (t TokenRequest) formData() url.Values {

	fData := make(url.Values)
	fData.Set("client_id", t.ClientID)
	fData.Set("access_type", t.AccessType)
	fData.Set("code", t.Code)
	fData.Set("grant_type", t.GrantType)
	fData.Set("redirect_uri", t.RedirectURI)

	return fData
}

func (t TokenRequest) post() (resp TokenResponse, err error) {
	httpResp, err := http.PostForm(tokenEndPoint, t.formData())
	if err != nil {
		return
	}

	if httpResp.StatusCode != http.StatusOK {
		//TODO make not suck with error struct
		bytes, _ := ioutil.ReadAll(httpResp.Body)
		err = fmt.Errorf("Token request responded in error %v: %v", httpResp.StatusCode, string(bytes))
		return
	}

	err = json.NewDecoder(httpResp.Body).Decode(&resp)

	return
}

func (t *TokenManager) GetInitialTokens(code string) error {

	req := TokenRequest{
		GrantType:   "authorization_code",
		Code:        code,
		AccessType:  "offline",
		ClientID:    t.ClientID + "@AMER.OAUTHAP",
		RedirectURI: "http://localhost",
	}

	res, err := req.post()
	if err != nil {
		return err
	}
	//TODO  keep an eye on this one, consider effects of overwritting tokens on error

	t.AccessToken, t.RefreshToken, err = res.Tokens()

	return err

}

// func (t TokenManager) RefreshToken() {

// 	reqBody, _ := TokenRequest{
// 		"grant_type":   "authorization_code",
// 		"access_type":  "offline",
// 		"code":         t.AccessToken.Code,
// 		"client_id":    t.ClientID,
// 		"redirect_uri": "http://localhost",
// 	}
// 	fmt.Print(string(reqBody))

// }

// take above example response and test it
// Deserialize json object using the json library and put into token response struct struct
// Ensure the tokens method converts the time correctly and test for edge cases. (ex, expires in 5 seconds)
// Ensure malform data works
// test worst case
//spit into files
// make tokens private
