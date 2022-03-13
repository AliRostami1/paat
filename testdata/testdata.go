package testdata

import (
	"fmt"
	"reflect"
)

var MapOfIntTestData map[string]int = map[string]int{}

func init() {
	for i := 1; i <= 10; i += 1 {
		MapOfIntTestData[fmt.Sprintf("field%d", i)] = i
	}
}

type StructTestDataType struct {
	Field1 string
	Field2 int
	Field3 float64
}

var StructTestData = StructTestDataType{
	Field1: "field1Value",
	Field2: 2,
	Field3: 3.0,
}

var ArrayOfStructsData = []StructTestDataType{}

func init() {
	for i := 1; i <= 10; i += 1 {
		ArrayOfStructsData = append(ArrayOfStructsData, StructTestDataType{
			Field1: fmt.Sprintf("field%dValue", i),
			Field2: i,
			Field3: float64(i),
		})
	}
}

var ArrayOfStructsDataReflect = []reflect.Value{}

var ArrayTestData []interface{} = []interface{}{}

func init() {
	for i := 1; i <= 10; i += 1 {
		ArrayTestData = append(ArrayTestData, StructTestDataType{
			Field1: fmt.Sprintf("field%dValue", i),
			Field2: i,
			Field3: float64(i),
		})
	}

	for i := 11; i <= 20; i += 1 {
		ArrayTestData = append(ArrayTestData, i)
	}
}

func init() {
	for i := 1; i <= 10; i += 1 {
		ArrayOfStructsDataReflect = append(ArrayOfStructsDataReflect, reflect.ValueOf(StructTestDataType{
			Field1: fmt.Sprintf("field%dValue", i),
			Field2: i,
			Field3: float64(i),
		}))
	}
}

type ComplexTestDataType struct {
	Field1 string
	Field2 string
	Field3 string
	Field4 int
	Field5 float64
	Field6 struct {
		EmbeddedField1 string
		EmbeddedField2 []string
		EmbeddedField3 string
	}
	Field7 []struct {
		EmbeddedArrayField1 string
		EmbeddedArrayField2 string
		EmbeddedArrayField3 string
	}
}

var ComplexTestData ComplexTestDataType = ComplexTestDataType{
	Field1: "Field1Value",
	Field2: "Field2Value",
	Field3: "Field3Value",
	Field4: 0,
	Field5: 0,
	Field6: struct {
		EmbeddedField1 string
		EmbeddedField2 []string
		EmbeddedField3 string
	}{
		EmbeddedField1: "EmbeddedField1Value",
		EmbeddedField2: []string{"EmbeddedField1ArrayValue", "EmbeddedField2ArrayValue", "EmbeddedField3ArrayValue"},
		EmbeddedField3: "EmbeddedField3Value",
	},
	Field7: []struct {
		EmbeddedArrayField1 string
		EmbeddedArrayField2 string
		EmbeddedArrayField3 string
	}{
		{
			EmbeddedArrayField1: "EmbeddedArrayField1Value",
			EmbeddedArrayField2: "EmbeddedArrayField2Value",
			EmbeddedArrayField3: "EmbeddedArrayField3Value",
		},
		{
			EmbeddedArrayField1: "EmbeddedArrayField1Value",
			EmbeddedArrayField2: "EmbeddedArrayField2Value",
			EmbeddedArrayField3: "EmbeddedArrayField3Value",
		},
		{
			EmbeddedArrayField1: "EmbeddedArrayField1Value",
			EmbeddedArrayField2: "EmbeddedArrayField2Value",
			EmbeddedArrayField3: "EmbeddedArrayField3Value",
		},
		{
			EmbeddedArrayField1: "EmbeddedArrayField1Value",
			EmbeddedArrayField2: "EmbeddedArrayField2Value",
			EmbeddedArrayField3: "EmbeddedArrayField3Value",
		},
		{
			EmbeddedArrayField1: "EmbeddedArrayField1Value",
			EmbeddedArrayField2: "EmbeddedArrayField2Value",
			EmbeddedArrayField3: "EmbeddedArrayField3Value",
		},
		{
			EmbeddedArrayField1: "EmbeddedArrayField1Value",
			EmbeddedArrayField2: "EmbeddedArrayField2Value",
			EmbeddedArrayField3: "EmbeddedArrayField3Value",
		},
		{
			EmbeddedArrayField1: "EmbeddedArrayField1Value",
			EmbeddedArrayField2: "EmbeddedArrayField2Value",
			EmbeddedArrayField3: "EmbeddedArrayField3Value",
		},
	},
}
