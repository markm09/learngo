// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/inancgumus/learngo/x-tba/2-methods/xxx-log-parser-methods/packaged/metrics"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	parser, report := metrics.NewParser(), metrics.NewReport()
	for in.Scan() {
		report.Update(parser.Parse(in.Text()))
	}
	s, _ := json.Marshal(report)
	fmt.Println(string(s))
	summarize(report, parser.Err(), in.Err())
}
