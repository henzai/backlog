package backlog

import (
	"fmt"
)

const (
	TextCustomField = iota + 1
	TextAreaCustomField
	NumericCustomField
	DateCustomField
	SingleListCustomField
	MultipleListCustomField
	CheckBoxCustomField
	RadioCustomField
)

type CustomField struct {
	ID          int
	FieldTypeID int
	Name        string
	Value       interface{}
}

func (c *CustomField) GetCustomFieldValue() {

	switch c.FieldTypeID {
	case TextCustomField:
		fmt.Println(TextCustomField)
	case TextAreaCustomField:
	case NumericCustomField:
	case DateCustomField:
	case SingleListCustomField:
	case MultipleListCustomField:
	case CheckBoxCustomField:
	case RadioCustomField:
	default:
	}
}
