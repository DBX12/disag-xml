package result

import (
	"encoding/xml"
)

type Shot struct {
	XMLName xml.Name `xml:"shot"`
	// CompetitionDateTime date and time of the competition
	CompetitionDateTime DisagDateTime `xml:"datetime,attr"`
	// ShotDateTime date and time of the shot
	ShotDateTime string `xml:"shotTime,attr"`
	// MenuId ID of the selected menu entry
	MenuId string `xml:"menuid,attr"`
	// Remark remark, e.g. penalties
	Remark string `xml:"remark,attr"`
	// CommentValid comments about validity of a shot, e.g. ricochet of another shooter
	CommentValid string `xml:"commentvalid,attr"`
	// Uuid unique identifier of a shot
	Uuid string `xml:"uuid,attr"`
	// UuidResult currently undocumented
	UuidResult string `xml:"uuidresult,attr"`
	// Valuation currently undocumented
	Valuation string `xml:"valuation,attr"`
	// DiskType type of the disk used for this shot
	DiskType string `xml:"disktyp,attr"`
	// WeaponsType currently undocumented
	WeaponsType string `xml:"weaponstype,attr"`
	// Source currently undocumented
	Source string `xml:"source,attr"`
	// DecimalResult result with decimal value
	DecimalResult SingleDecimalValue `xml:"dec,attr"`
	// FullRings result in full rings
	FullRings int `xml:"full,attr"`
	// DecimalPart decimal part of the result
	DecimalPart int `xml:"onlydec,attr"`
	// X x coordinate of the shot on the target surface
	X int `xml:"x,attr"`
	// Y y coordinate of the shot on the target surface
	Y int `xml:"y,attr"`
	// DistanceToCenter distance to the center of the disk in 1/100 millimeters
	DistanceToCenter SingleDecimalValue `xml:"teiler,attr"`
	// TrafficLightStatus status of the traffic light during this shot
	TrafficLightStatus string `xml:"tlstatus,attr"`
	// TlTime currently undocumented
	TlTime string `xml:"tltime,attr"`
	// Tlight currently undocumented
	Tlight string `xml:"tlight,attr"`
	// ShootingRange number of the shooting range
	ShootingRange int `xml:"shootingrange,attr"`
	// ShootersId start number of the shooter
	ShootersId int `xml:"shootersid,attr"`
	// CompetitionId currently undocumented
	CompetitionId int `xml:"competitionid,attr"`
	// Shot number of the shot (zero based)
	Shot int `xml:"shot,attr"`
	// ShotCount number of the shot (zero based)
	ShotCount int `xml:"shotcount,attr"`
	// Run currently undocumented
	Run string `xml:"run,attr"`
	// Angle currently undocumented
	Angle int `xml:"winkel,attr"`
	// InnerTen currently undocumented
	InnerTen string `xml:"InnerTen,attr"`
	// IsInnerTen if this shot is fully within the ten circle
	IsInnerTen bool `xml:"isinnerten,attr"`
	// IsValid is it a valid shot
	IsValid bool `xml:"isvalid,attr"`
	// IsHot shot made for a competition
	IsHot bool `xml:"ishot,attr"`
	// IsWarmup shot made during warmup (shot is not valid for competition)
	IsWarmup bool `xml:"iswarm,attr"`
	// IsDummy shot was generated
	IsDummy bool `xml:"dummy,attr"`
	// Placeholder currently undocumented
	Placeholder bool `xml:"placeholder,attr"`
	ShotValue   int  `xml:",chardata"`
}
