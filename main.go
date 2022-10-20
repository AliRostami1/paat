package main

import (
	"fmt"
	"log"

	"github.com/AliRostami1/paat/pkg/parser"
	testdata "github.com/AliRostami1/paat/testdata/go"
)

func main() {

	c, err := parser.Parse(testdata.ComplexTestData)
	if err != nil {
		log.Fatalf("parser: %v", err)
	}
	canvas := c.Draw()
	fmt.Print(canvas.String())

	c, err = parser.Parse("testsds\ntest\ntest\ntest\ntest")
	if err != nil {
		log.Fatalf("parser: %v", err)
	}
	canvas = c.Draw()
	fmt.Print(canvas.String())
}
