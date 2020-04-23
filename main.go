package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/HannanSolo/go-td-ameritrade-api/auth"
)

func main() {
	//TODO make conditional if we have a valid refresh token
	const clientID = "997F0OYWTJXUZZNG9HSNCFIAUHMGYBQW"
	//obtain a initial auth	 token by auth through browser

	reader := bufio.NewReader(os.Stdin)

	login := auth.Login{ClientID: clientID, RedirectURI: "http://localhost"}

	fmt.Printf("Put this URL into your browser to login\n %v\n", login.URL())

	fmt.Print("Please paste the auth url you got: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	u, err := url.Parse(strings.TrimSpace(text))
	if err != nil {
		log.Fatalf("Error parsing url: %v", err)
	}

	code := login.ExtractCode(u)
	fmt.Println("\n\n\n" + code)
	if code == "" {
		log.Fatal("Url did not contain a code parameter")
	}

	han := auth.TokenManager{ClientID: clientID}

	err = han.GetInitialTokens(code)
	if err != nil {
		log.Fatalf("Error getting tokens %v", err)
	}

	fmt.Printf("%#v\n", han)

	// get a initial refresh token.

	// - contruct post request with body parameters
	// - Send post request to abtain refresh and auth token
	// - Parse response into TokenResponse type
	// - Validate contents
	// - calculate absolute expire time

}

//

//https://auth.tdameritrade.com/oauth?client_id=997F0OYWTJXUZZNG9HSNCFIAUHMGYBQW%40AMER.OAUTHAP&response_type=code&redirect_uri=http%3A%2F%2Flocalhost
//https://auth.tdameritrade.com/oauth?client_id=997F0OYWTJXUZZNG9HSNCFIAUHMGYBQW%40AMER.OAUTHAP&redirect_uri=http%3A%2F%2Flocalhost&response_type=code
//Create Http handler that respond to the redirect uri

//pickup code that returns auth token
// |
// v
//https://localhost/?code=EGyL55IV%2F8upykp35KbI%2FOurRie5KqM571xxKiREopMgwdfIOrQPHtysF9jelFqiuIqdHOIRdK%2Fbb37jNYmwTwSYgdzVwvYeuP5zjfB7g2VHuzJPA7bFYKnCAFvhmIy5d0YTZuvfDcH1BN5AZbtAm8U03nIvZSl84fi4R3ur6j0jy5BPCe9EWyp9RUM78hYpLx8rBScu9rSHOMhxATctN5O3mIr6ZCc8VDSVpHTsVPYYYRjkjlb%2F4heHHUxnPcCYPuO71cTo6W7CHoEXEpGG%2FtQltldsurUsSvKb8rAyVdhTtI1Fg5MheJIZYNcRgMnD6fy1rzq6T4nqEDaevnWR6teQ9%2BbE20OLpAMUhw7yOXAaZzmwV%2Biz5hnMD7iKDmWUJZS2rixY5tHTr95FmCrZ%2BxkP%2BXWJggCD6rCTlZSnbYg1ourjI9IFGJkOrBO100MQuG4LYrgoVi%2FJHHvl2uQwIZ9AnRCzt0BwcqBpjHgrGF0mBRDiVd1HpWQ3Al2kveB6l7zRr6CbuhM4A66QvdniXtmNJsc57a9fvtvdPIWPBdo0iNR1XlhcmNF6Zzbvqk90uqEhola0uBbypWp23kNRLV3mkOE3%2FO77q%2FCPcFN0ryF%2FNSuMtIeJaWO38B%2FMgQeqm3chpxFbPy0dmDkGOR2temhmrRptzcCwq5rhkr8PiyVO0HAcosSnPaTxm1hAycDA1G2HQClt1bz5ypu1mj%2FHzSacioE52AUzTVvYcTmTKkF1EeotizC%2FcRPjqp8BxQK39ft0btvc%2FCdKnQQdn7NgqDuSyd83FzgQim1scjmg5ErewCsPJXCwejxeheG69dfIMjDcI4i%2FOsoBOc2Jf9%2BjMJkdlR2nGXBk1V6e6ryc%2B7%2BmQxfwMR5qNOSjWpO0xWQd1WccXfEzvO4%3D212FD3x19z9sWBHDJACbC00B75E

//setup server that serves the http handler on a known port ex(8080)

//recieve code and put into file to save
//os package to open file and create encoder on the file to create json
//create write and load functions
