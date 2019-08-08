package main

import "github.com/gojisvm/gojis/internal/parser"

func main() {
	p := parser.New()
	err := p.ParseFile("testdata/small.js")
	_ = err
}
