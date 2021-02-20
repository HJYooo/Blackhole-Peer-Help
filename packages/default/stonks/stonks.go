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
	Ma