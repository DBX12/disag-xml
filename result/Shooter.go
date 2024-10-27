package result

import "encoding/xml"

type Shooter struct {
	XMLName                             xml.Name           `xml:"shooter"`
	Identification                      int                `xml:"identification,attr"`
	BirthDateIso                        DisagDateIso       `xml:"birthdateiso,attr"`
	IdShooters                          int                `xml:"idShooters,attr"`
	FirstName                           string             `xml:"firstname,attr"`
	LastName                            string             `xml:"lastname,attr"`
	FullName                            string             `xml:"fullname,attr"`
	ClubName                            string             `xml:"clubsname,attr"`
	FactorClass                         SingleDecimalValue `xml:"factor_class,attr"`
	SeriesFactorClass                   SingleDecimalValue `xml:"series_factor_class,attr"`
	Factor                              SingleDecimalValue `xml:"factor,attr"`
	FactorWeapon                        SingleDecimalValue `xml:"factor_weapon,attr"`
	SeriesFactor                        SingleDecimalValue `xml:"series_factor,attr"`
	SeriesFactorWeapon                  SingleDecimalValue `xml:"series_factor_weapon,attr"`
	LastDistance                        SingleDecimalValue `xml:"lastdistance,attr"`
	LastShotCounts                      SingleDecimalValue `xml:"lastshotcounts,attr"`
	DistancesStart                      SingleDecimalValue `xml:"distances_start,attr"`
	TotalScoreAvg                       SingleDecimalValue `xml:"totalscore_avg,attr"`
	TotalScoreAvgDecimals               SingleDecimalValue `xml:"totalscore_avg_dec,attr"`
	TotalScore                          int                `xml:"totalscore,attr"`
	TotalScoreD                         SingleDecimalValue `xml:"totalscore_d,attr"`
	TotalScoreOnlyD                     int                `xml:"totalscore_onlyd,attr"`
	CountX                              int                `xml:"count_X,attr"`
	BestDistanceToCenter                SingleDecimalValue `xml:"bester_teiler,attr"`
	BestDistanceToCenterFromSum         SingleDecimalValue `xml:"bester_teiler_aus_summe,attr"`
	BestDistanceToCenter2               SingleDecimalValue `xml:"bester_teiler_2,attr"`
	BestDistanceToCenter3               SingleDecimalValue `xml:"bester_teiler_3,attr"`
	OriginalBestDistanceToCenterFromSum SingleDecimalValue `xml:"bester_teiler_aus_summe_orig,attr"`
	OriginalBestDistanceToCenter        SingleDecimalValue `xml:"orig_teiler,attr"`
	OriginalBestDistanceToCenter2       SingleDecimalValue `xml:"bester_teiler_2_orig,attr"`
	OriginalBestDistanceToCenter3       SingleDecimalValue `xml:"bester_teiler_3_orig,attr"`
	InTeam                              bool               `xml:"inteam,attr"`
	Shots                               Shots              `xml:"shots"`
}
