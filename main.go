package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/AliRostami1/tabler/pkg/parser"
	"golang.org/x/term"
)

func main() {
	file, err := os.Open("testdata/random.json")
	if err != nil {
		log.Fatalf("opening file: %v", err)
	}
	defer file.Close()

	firstLetter := make([]byte, 1)

	_, err = file.Read(firstLetter)
	file.Seek(0, 0)
	if err != nil {
		log.Fatalf("can't read from the file: %v", err)
	}
	jsonDecoder := json.NewDecoder(file)

	var cell *parser.Cell

	switch firstLetter[0] {
	case '[':
		unknwonArray := []interface{}{}
		err = jsonDecoder.Decode(&unknwonArray)
		if err != nil {
			log.Fatalf("decoding json from file: %v", err)
		}

		cell, err = parser.Parse(unknwonArray)
		if err != nil {
			log.Fatalf("parser failed: %v", err)
		}

	case '{':
		unknwonStruct := map[string]interface{}{}
		err = jsonDecoder.Decode(&unknwonStruct)
		if err != nil {
			log.Fatalf("decoding json from file: %v", err)
		}

		cell, err = parser.Parse(unknwonStruct)
		if err != nil {
			log.Fatalf("parser failed: %v", err)
		}

	default:
		log.Fatal("invalid json")
	}

	ttyWidth, ttyHeight, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("getting terminal dimensions: %v", err)
	}

	log.Printf("terminal width=%d, height=%d", ttyWidth, ttyHeight)

	// if cell.BorderBoxWidth() > ttyWidth {
	// 	cell.SetBorderBoxWidth(ttyWidth - 1)
	// }

	canvas := cell.Draw()
	fmt.Print(canvas.String())
}
