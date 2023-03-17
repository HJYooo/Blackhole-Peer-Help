
package main

import (
	"bytes"
	"html/template"
	"strconv"
)

var fmap template.FuncMap

func init() {
	fmap = template.FuncMap{}
	fmap["comma"] = func(n int) string {
		in := strconv.FormatInt(int64(n), 10)
		numOfDigits := len(in)
		if n < 0 {
			numOfDigits-- // First character is the - sign (not a digit)
		}
		numOfCommas := (numOfDigits - 1) / 3

		out := make([]byte, len(in)+numOfCommas)
		if n < 0 {
			in, out[0] = in[1:], '-'
		}