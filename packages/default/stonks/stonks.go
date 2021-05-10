package main

import (
	"time"

	"github.com/jonfriesen/finance-go"
	"github.com/jonfriesen/finance-go/chart"
	"github.com/jonfriesen/finance-go/datetime"
	"github.com/jonfriesen/finance-go/quote"
	"github.com/pkg/errors"
)

type MarketState string

type YQuote struct {
	Chart           string
	Trend           string // can be bigup (>3%), up, drop or bigdrop (<3%)
	Symbol          string
	Name            string
	Currency        string
	MarketState     string
	MarketPrice     float64
	MarketChange    float64
	MarketChangePct float64
}

func Main(args map[string]interface{}) map[string]interface{} {
	symbol, ok := args["symbol"].(string)
	if !ok {
		symbol = "DOCN"
	}

	resp := make(map[string]interface{})

	q := getQuote(symbol)

	// b, err := json.Marshal(q)
	// if err != nil {
	// 	resp["error"] = errors.Wrap(err, "marshalling quote object")
	// 	return resp
	// }

	start := time.Now().Add(time.Hour * 24 * 30 * -1)
	end := time.Now()

	history := 