
package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/jonfriesen/finance-go"
	"github.com/pkg/errors"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

func Sparkline(symbol string, charBars []*finance.ChartBar) (*bytes.Buffer, error) {
	var dates []time.Time
	var yv []float64
	for _, ts := range charBars {
		parsed := time.Unix(int64(ts.Timestamp), 0)
		dates = append(dates, parsed)

		close, _ := ts.Close.Float64()
		yv = append(yv, close)
	}

	priceSeries := chart.TimeSeries{
		XValues: dates,
		YValues: yv,
	}
