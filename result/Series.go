package result

import (
	"encoding/xml"
	"fmt"
)

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

func (s Series) SumFullRings(start int, end int) (int, error) {
	err := s.validateStartAndEnd(start, end)
	if err != nil {
		return 0, err
	}

	sum := 0
	for i := start; i <= end; i++ {
		sum += s.Shots[i].FullRings
	}
	return sum, nil
}
func (s Series) SumDecimalResult(start int, end int) (SingleDecimalValue, error) {
	err := s.validateStartAndEnd(start, end)
	if err != nil {
		return 0, err
	}

	sum := SingleDecimalValue(0)
	for i := start; i <= end; i++ {
		sum += s.Shots[i].DecimalResult
	}
	return sum, nil
}

func (s Series) validateStartAndEnd(start, end int) error {
	if start < 0 {
		return fmt.Errorf("invalid parameter: start must be greater than zero")
	}
	if start > end {
		return fmt.Errorf("invalid parameter: start must be less than end")
	}
	if end >= len(s.Shots) {
		return fmt.Errorf("invalid parameter: series has less shots than requested with end")
	}
	return nil
}
