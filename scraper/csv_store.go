package scraper

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

type CandleWriter struct {
	*csv.Writer
}

// create a method that takes in a candle and returns it as a csv write friendly slice
func (w CandleWriter) Write(candle Candle) error {

	return w.Writer.Write([]string{
		candle.Open.String(),
		candle.High.String(),
		candle.Low.String(),
		candle.Close.String(),
		strconv.Itoa(candle.Volume),
		strconv.FormatInt(candle.Datetime, 10),
	})

}

//get or create a csv writer based on a ticker csv file
func NewCandleWriter(w io.Writer) CandleWriter {
	return CandleWriter{csv.NewWriter(w)}

}

//create a  ticker to file and return an io writer
func TickerToFile(ticker string) (*os.File, error) {
	//TODO have a scenario  where we check the first line to compare it to a header or a data point, respond accordingly.

	//open file, if not exist then create it
	f, err := os.OpenFile(ticker+".csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		f, err = os.Create(ticker + ".csv")
		if err != nil {
			return nil, err
		}
		//debating on adding header row when creating a file
	}

	return f, err
}

//return err on create
//TODO wrap errors
