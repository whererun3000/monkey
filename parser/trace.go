package parser

import (
	"fmt"
	"strings"
)

func trace(msg string) string {
	tracePrint("BEGIN " + msg)
	incIdent()
	return msg
}

func untrace(msg string) {
	decIdent()
	tracePrint("END " + msg)
}

func tracePrint(msg string) {
	fmt.Printf("%s%s\n", strings.Repeat("\t", traceLevel), msg)
}

var traceLevel int

func incIdent() { traceLevel++ }
func decIdent() { traceLevel-- }
