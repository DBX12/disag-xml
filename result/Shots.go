package result

import "encoding/xml"

type Shots struct {
	XMLName xml.Name `xml:"shots"`
	Series  []Series `xml:"series"`
}
