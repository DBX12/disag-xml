package result

import "encoding/xml"

type Result struct {
	XMLName                xml.Name      `xml:"result"`
	Type                   string        `xml:"type,attr"`
	FlexVersion            string        `xml:"flexVersion,attr"`
	Date                   DisagDateTime `xml:"date,attr"`
	Location               string        `xml:"location,attr"`
	Name                   string        `xml:"name,attr"`
	Id                     int           `xml:"id,attr"`
	FlexUuid               string        `xml:"flexUUID,attr"`
	PrintDateTime          DisagDateTime `xml:"printDateTime,attr"`
	PrintDateTimeInvariant string        `xml:"printDateTimeInvariant,attr"`
	DelayInSeconds         int           `xml:"delayInSeconds,attr"`
	HideBestResults        bool          `xml:"hideBestResults,attr"`
	HideBestResultsCount   int           `xml:"hideBestResultsCount,attr"`
	Shooters               ShootersList  `xml:"shooters"`
}
