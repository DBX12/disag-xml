package result

import (
	"encoding/xml"
	"time"
)

// DisagDateIso represents a date in the format YYYYMMDD
type DisagDateIso time.Time

func (d *DisagDateIso) UnmarshalXMLAttr(attr xml.Attr) error {
	v, err := d.parseString(attr.Value)
	if err != nil {
		return err
	}
	*d = DisagDateIso(*v)
	return nil
}

func (d *DisagDateIso) parseString(str string) (*time.Time, error) {
	if str == "" {
		return &time.Time{}, nil
	}

	t, err := time.Parse("20060102", str)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
