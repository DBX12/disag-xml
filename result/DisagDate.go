package result

import (
	"encoding/xml"
	"time"
)

type DisagDateTime time.Time

func (d *DisagDateTime) UnmarshalXMLAttr(attr xml.Attr) error {
	return d.parseString(attr.Value)
}

func (d *DisagDateTime) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	//29.11.2023 20:44:39
	var str string

	err := dec.DecodeElement(&str, &start)
	if err != nil {
		return err
	}
	return d.parseString(str)
}

func (d *DisagDateTime) parseString(str string) error {
	t, err := time.Parse("02.01.2006 15:04:05", str)
	if err != nil {
		return err
	}
	*d = DisagDateTime(t)
	return nil
}
