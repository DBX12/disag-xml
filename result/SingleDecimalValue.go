package result

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// SingleDecimalValue values of this type should be represented with only a single decimal
type SingleDecimalValue float32

func (d *SingleDecimalValue) String() string {
	return regexp.MustCompile(`\d+\.\d`).FindString(fmt.Sprintf("%f", *d))
}

func (d *SingleDecimalValue) UnmarshalXMLAttr(attr xml.Attr) error {
	return d.parseString(attr.Value)
}

func (d *SingleDecimalValue) parseString(str string) error {
	str = strings.Replace(str, ",", ".", 1)
	flt, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return err
	}
	*d = SingleDecimalValue(flt)
	return nil
}
