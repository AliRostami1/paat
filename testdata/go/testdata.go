package testdata

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
