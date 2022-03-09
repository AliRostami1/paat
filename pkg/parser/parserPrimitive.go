package parser

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func (c *Cell) parsePrimitive(in interface{}) error {
	// single cell representing the primitive value

	// it's a primitive so the type of the parent Cell whould be
	// PrimitiveCell which indicates there is a string in it's
	// contetnt field
	c.Type = PrimitiveCell
	// as this is a PrimitiveCell, we extract the string
	// reperesentation of it's value and assign it to it's
	// content field
	strRep := fmt.Sprint(in)
	c.Content = strRep
	// width of the cell will be the number of Unicode code points
	// in the string reperesntation of it
	c.Width = utf8.RuneCountInString(strRep)
	// height of the cell will be number of \n(newline characters)
	// in the string reperesntation of it
	c.Height = strings.Count(strRep, "\n") + 1
	return nil
}
