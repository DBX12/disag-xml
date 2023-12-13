package result

import "encoding/xml"

type Series struct {
	XMLName xml.Name `xml:"series"`
	// DsbPunkteS currently undocumented
	DsbPunkteS SingleDecimalValue `xml:"dsbpunkte_s,attr"`
	// TotalScoreOnlyD currently undocumented - reverse engineered to be the sum of the  Shot.DecimalPart field of Shots
	TotalScoreOnlyD int `xml:"totalscore_onlyd,attr"`
	// TotalScore currently undocumented - reverse engineered to be the sum of the  Shot.FullRings field of Shots
	TotalScore int `xml:"totalscore,attr"`
	// TotalScoreD currently undocumented - reverse engineered to be the sum of the Shot.DecimalResult field of Shots
	TotalScoreD SingleDecimalValue `xml:"totalscore_d,attr"`
	// BestDistanceToCenter currently undocumented - smallest value of the Shot.DistanceToCenter field of Shots
	BestDistanceToCenter SingleDecimalValue `xml:"bester_teiler,attr"`
	// OuterShooterTotalScoreBestCount currently undocumented
	OuterShooterTotalScoreBestCount int `xml:"outershooterstotalscore_bestcount,attr"`
	// OuterShooterTotalScoreBestCountD currently undocumented
	OuterShooterTotalScoreBestCountD int `xml:"outershooterstotalscore_bestcount_d,attr"`
	// OuterShooterTotalScoreBestCountOnlyD currently undocumented
	OuterShooterTotalScoreBestCountOnlyD int `xml:"outershooterstotalscore_bestcount_onlyd,attr"`
	// Shots shots in this series
	Shots []Shot `xml:"shot"`
}

func (s Series) SumFullRings(start int, end int) int {
	sum := 0
	for i := start; i < end; i++ {
		sum += s.Shots[i].FullRings
	}
	return sum
}

func (s Series) SumDecimalResult(start int, end int) SingleDecimalValue {
	sum := SingleDecimalValue(0)
	for i := 0; i < end; i++ {
		sum += s.Shots[i].DecimalResult
	}
	return sum
}
