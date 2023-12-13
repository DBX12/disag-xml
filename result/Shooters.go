package result

import "encoding/xml"

type ShootersList struct {
	XMLName  xml.Name  `xml:"shooters"`
	Title    string    `xml:"title,attr"`
	Uuid     string    `xml:"uuid,attr"`
	Shooters []Shooter `xml:"shooter"`
}
