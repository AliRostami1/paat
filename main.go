package main

import (
	"fmt"
	"log"

	"github.com/AliRostami1/tabler/pkg/parser"
)

type TestData struct {
	Field1 string
	Field2 string
	Field3 string
}

func main() {
	c, err := parser.Parse([]TestData{
		{
			Field1: "hey",
			Field2: "hoy",
			Field3: "damn",
		},
		{
			Field1: "hey",
			Field2: "hoy",
			Field3: "damn",
		},
		{
			Field1: "hey",
			Field2: "hoy",
			Field3: "damn",
		},
		{
			Field1: "hey",
			Field2: "hoy",
			Field3: "damn",
		},
	})
	if err != nil {
		log.Fatalf("parser: %v", err)
	}
	canvas := c.Draw()
	fmt.Print(canvas.String())

	c, err = parser.Parse("testsds")
	if err != nil {
		log.Fatalf("parser: %v", err)
	}
	canvas = c.Draw()
	fmt.Print(canvas.String())
}
