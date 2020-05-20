package scraper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const baseURL = "https://api.tdameritrade.com/v1/marketdata/"

const (
	Day     = periodType("day")
	Month   = periodType("month")
	Year    = periodType("year")
	YTD     = periodType("ytd")
	Daily   = frequencyType("daily")
	Minute  = frequencyType("minute")
	Weekly  = frequencyType("weekly")
	Monthly = frequencyType("monthly")
)

type periodType string
type frequencyType string

//This conforms to the https://developer.tdameritrade.com/price-history/apis/get/marketdata/%7Bsymbol%7D/pricehistory
type Request struct {
	Ticker            string
	PeriodType        periodType
	Period            int
	FrequencyType     frequencyType
	Frequency         int
	EndDate           time.Time
	StartDate         time.Time
	ExtendedHoursData bool
}

type Candle struct {
	Close    json.Number `json:"close"`
	Open     json.Number `json:"open"`
	Datetime int64       `json:"datetime"`
	High     json.Number `json:"high"`
	Low      json.Number `json:"low"`
	Volume   int         `json:"volume"`
}

type Response struct {
	Candles Candle `json:"candles"`
	Empty   bool   `json:"empty"`
	Symbol  string `json:"symbol"`
}

func timeToUnixMsec(t time.Time) int64 {

	nsec := time.Duration(t.UnixNano()) * time.Nanosecond

	return nsec.Milliseconds()
}

func (r *Request) HttpRequest() (*http.Request, error) {

	params := make(url.Values)

	params.Set("periodType", string(r.PeriodType))
	if r.Period != 0 {
		params.Set("period", fmt.Sprint(r.Period))
	}
	params.Set("frequencyType", string(r.FrequencyType))
	if r.Frequency != 0 {
		params.Set("frequency", fmt.Sprint(r.Frequency))
	}
	if !r.EndDate.IsZero() {
		params.Set("endDate", fmt.Sprint(timeToUnixMsec(r.EndDate)))
	}
	if !r.StartDate.IsZero() {
		params.Set("startdate", fmt.Sprint(timeToUnixMsec(r.StartDate)))
	}
	params.Set("needExtendedHoursData", fmt.Sprint(r.ExtendedHoursData))

	return http.NewRequest("GET", fmt.Sprintf("%s/%s/pricehistory?%s", baseURL, r.Ticker, params.Encode()), nil)
}
