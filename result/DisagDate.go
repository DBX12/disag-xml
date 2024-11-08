package result

import (
	"encoding/xml"
	"time"
)

type DisagDateTime time.Time

func (d *DisagDateTime) UnmarshalXMLAttr(attr xml.Attr) error {
	v, err := d.parseString(attr.Value)
	if err != nil {
		return err
	}
	*d = DisagDateTime(*v)
	return nil
}

func (d *DisagDateTime) parseString(str string) (*time.Time, error) {
	if str == "" {
		return &time.Time{}, nil
	}

	t, err := time.Parse("02.01.2006 15:04:05", str)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
