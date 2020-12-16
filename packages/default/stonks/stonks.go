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
	Tr