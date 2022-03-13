package main

import (
	"fmt"
	"log"

	"github.com/AliRostami1/tabler/pkg/parser"
	"github.com/AliRostami1/tabler/testdata"
)

func main() {

	c, err := parser.Parse(testdata.ComplexTestData)
	if err != nil {
		log.Fatalf("parser: %v", err)
	}
	canvas := c.Draw()
	fmt.Print(canvas.String())

	c, err = parser.Parse("testsds\ntest\ntest\ntest\ntest")
	// fmt.Printf("%#+v", *c)
	if err != nil {
		log.Fatalf("parser: %v", err)
	}
	canvas = c.Draw()
	fmt.Print(canvas.String())
}
