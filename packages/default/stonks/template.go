
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

		for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
			out[j] = in[i]
			if i == 0 {
				return string(out)
			}
			if k++; k == 3 {
				j, k = j-1, 0
				out[j] = ','
			}
		}
	}
	fmap["safeHTML"] = func(s string) template.HTML { return template.HTML(s) }
}

const quoteTmpl = `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
</head>
<body>
<h1><b>{{.Symbol}}</b> {{.Name}}</h1>
{{safeHTML .Chart}}