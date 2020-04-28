package auth

import (
	"net/url"
)

const loginEndPoint = "https://auth.tdameritrade.com/auth"

// Login holds our TD user's info
type Login struct {
	ClientID    string
	RedirectURI string
}

// URL creates the auth url for the user to login to their account
func (c Login) URL() *url.URL {
	u, err := url.Parse(loginEndPoint)
	if err != nil {
		panic("Url didnt parse")
	}

	//adding query values
	q := make(url.Values)
	q.Set("client_id", c.ClientID+"@AMER.OAUTHAP")
	q.Set("response_type", "code")
	q.Set("redirect_uri", c.RedirectURI)
	u.RawQuery = q.Encode()

	return u
}

// ExtractCode takes in a url and query it for the code value
func (l Login) ExtractCode(u *url.URL) string {
	return u.Query().Get("code")
}

// post https://api.tdameritrade.com/v1/oauth2/token
