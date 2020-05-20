package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/HannanSolo/go-td-ameritrade-api/auth"
	"github.com/HannanSolo/go-td-ameritrade-api/scraper"
)

func main() {
	//TODO make conditional if we have a valid refresh token
	const clientID = "1FSINCAKAGNKFIY9HXI4V56GQVPHAALZ"

	//obtain a initial auth	 token by auth through browser

	han := auth.TokenManager{ClientID: clientID}

	const file = "cred.json"
	if _, err := os.Stat(file); err == nil {

		err = han.Load(file)
		if err != nil {
			log.Fatalf("Was not able to read cred file %v", err)
		}
	}

	if han.LoginRequired() {
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

		err = han.GetInitialTokens(code)
		if err != nil {
			log.Fatalf("Error getting tokens %v", err)
		}

		han.Save(file)
	}

	if han.RefreshRequired() {
		err := han.RefreshTokens()
		if err != nil {
			log.Fatalf("Error refreshing tokens %v", err)
		}
		han.Save(file)
	}
	//fmt.Printf("%v\n", han)

	day1 := time.Unix(0, 1589290500000*int64(time.Millisecond)) //time.Date(2018, time.April, 20, 9, 30, 0, 0, time.Local)
	day2 := time.Unix(0, 1589827680000*int64(time.Millisecond)) //time.Date(2018, time.April, 30, 16, 30, 0, 0, time.Local)
	//create a request struct
	amddata := scraper.Request{Ticker: "AMD", FrequencyType: scraper.Minute, Frequency: 1, EndDate: day2, StartDate: day1, ExtendedHoursData: false}

	req, _ := amddata.HttpRequest()

	han.AddAuthorization(req)
	//now it has everything we need

	fmt.Println(req.URL.String())

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Print(string(body))
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
