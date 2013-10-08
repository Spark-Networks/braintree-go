package braintree

import (
	"encoding/xml"
)

type Search struct {
	XMLName string `xml:"search"`
	Fields  []interface{}
}

type TextField struct {
	XMLName    xml.Name
	Is         string `xml:"is,omitempty"`
	IsNot      string `xml:"is-not,omitempty"`
	StartsWith string `xml:"starts-with,omitempty"`
	EndsWidth  string `xml:"ends-with,omitempty"`
	Contains   string `xml:"contains,omitempty"`
}

type RangeField struct {
	XMLName xml.Name
	Is      float64 `xml:"is,omitempty"`
	Min     float64 `xml:"min,omitempty"`
	Max     float64 `xml:"max,omitempty"`
}

type MultiField struct {
	XMLName xml.Name
	Type    string   `xml:"type,attr"` // type=array
	Item    []string `xml:"item"`
}

func (s *Search) AddTextField(field string) *TextField {
	f := &TextField{XMLName: xml.Name{Local: field}}
	s.Fields = append(s.Fields, f)
	return f
}

func (s *Search) AddRangeField(field string) *RangeField {
	f := &RangeField{XMLName: xml.Name{Local: field}}
	s.Fields = append(s.Fields, f)
	return f
}

func (s *Search) AddMultiField(field string) *MultiField {
	f := &MultiField{
		XMLName: xml.Name{Local: field},
		Type:    "array",
	}
	s.Fields = append(s.Fields, f)
	return f
}

func (s *Search) ToXML() string {
	b, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (f *MultiField) AddItem(s string) {
	f.Item = append(f.Item, s)
}
