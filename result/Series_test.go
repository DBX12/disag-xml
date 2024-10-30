package result

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getShotSeries() []Shot {
	return []Shot{
		{
			FullRings:     8,
			DecimalResult: 8.2,
		},
		{
			FullRings:     2,
			DecimalResult: 2.5,
		},
		{
			FullRings:     1,
			DecimalResult: 1.4,
		},
		{
			FullRings:     9,
			DecimalResult: 9.4,
		},
		{
			FullRings:     5,
			DecimalResult: 5.7,
		},
	}
}

func TestSeries_SumFullRings(t *testing.T) {
	type fields struct {
		XMLName                              xml.Name
		DsbPunkteS                           SingleDecimalValue
		TotalScoreOnlyD                      int
		TotalScore                           int
		TotalScoreD                          SingleDecimalValue
		BestDistanceToCenter                 SingleDecimalValue
		OuterShooterTotalScoreBestCount      int
		OuterShooterTotalScoreBestCountD     int
		OuterShooterTotalScoreBestCountOnlyD int
		Shots                                []Shot
	}
	type args struct {
		start int
		end   int
	}

	shots := getShotSeries()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "Validation error",
			fields:  fields{Shots: shots},
			args:    args{start: -1},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "Good case",
			fields:  fields{Shots: shots},
			args:    args{3, 4},
			want:    14,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Series{
				XMLName:                              tt.fields.XMLName,
				DsbPunkteS:                           tt.fields.DsbPunkteS,
				TotalScoreOnlyD:                      tt.fields.TotalScoreOnlyD,
				TotalScore:                           tt.fields.TotalScore,
				TotalScoreD:                          tt.fields.TotalScoreD,
				BestDistanceToCenter:                 tt.fields.BestDistanceToCenter,
				OuterShooterTotalScoreBestCount:      tt.fields.OuterShooterTotalScoreBestCount,
				OuterShooterTotalScoreBestCountD:     tt.fields.OuterShooterTotalScoreBestCountD,
				OuterShooterTotalScoreBestCountOnlyD: tt.fields.OuterShooterTotalScoreBestCountOnlyD,
				Shots:                                tt.fields.Shots,
			}
			got, err := s.SumFullRings(tt.args.start, tt.args.end)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSeries_SumDecimalResult(t *testing.T) {
	type fields struct {
		XMLName                              xml.Name
		DsbPunkteS                           SingleDecimalValue
		TotalScoreOnlyD                      int
		TotalScore                           int
		TotalScoreD                          SingleDecimalValue
		BestDistanceToCenter                 SingleDecimalValue
		OuterShooterTotalScoreBestCount      int
		OuterShooterTotalScoreBestCountD     int
		OuterShooterTotalScoreBestCountOnlyD int
		Shots                                []Shot
	}
	type args struct {
		start int
		end   int
	}

	shots := getShotSeries()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    SingleDecimalValue
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "Validation error",
			fields:  fields{Shots: shots},
			args:    args{start: -1},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:   "Good case",
			fields: fields{Shots: shots},
			args:   args{0, 2},
			// this is needed to trick the compiler to perform the calculation at runtime
			// since the compiler-time and runtime results differ by 0.01. Floating points, yay.
			want:    SingleDecimalValue(float32(8.2) + float32(2.5) + float32(1.4)),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Series{
				XMLName:                              tt.fields.XMLName,
				DsbPunkteS:                           tt.fields.DsbPunkteS,
				TotalScoreOnlyD:                      tt.fields.TotalScoreOnlyD,
				TotalScore:                           tt.fields.TotalScore,
				TotalScoreD:                          tt.fields.TotalScoreD,
				BestDistanceToCenter:                 tt.fields.BestDistanceToCenter,
				OuterShooterTotalScoreBestCount:      tt.fields.OuterShooterTotalScoreBestCount,
				OuterShooterTotalScoreBestCountD:     tt.fields.OuterShooterTotalScoreBestCountD,
				OuterShooterTotalScoreBestCountOnlyD: tt.fields.OuterShooterTotalScoreBestCountOnlyD,
				Shots:                                tt.fields.Shots,
			}
			got, err := s.SumDecimalResult(tt.args.start, tt.args.end)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSeries_validateStartAndEnd(t *testing.T) {
	type fields struct {
		XMLName                              xml.Name
		DsbPunkteS                           SingleDecimalValue
		TotalScoreOnlyD                      int
		TotalScore                           int
		TotalScoreD                          SingleDecimalValue
		BestDistanceToCenter                 SingleDecimalValue
		OuterShooterTotalScoreBestCount      int
		OuterShooterTotalScoreBestCountD     int
		OuterShooterTotalScoreBestCountOnlyD int
		Shots                                []Shot
	}
	type args struct {
		start int
		end   int
	}

	shots := getShotSeries()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "Invalid parameter: start less than zero",
			fields:  fields{Shots: shots},
			args:    args{-2, 3},
			wantErr: assert.Error,
		},
		{
			name:    "Invalid parameter: start greater than end",
			fields:  fields{Shots: shots},
			args:    args{5, 3},
			wantErr: assert.Error,
		},
		{
			name:    "Invalid parameter: end greater than series size",
			fields:  fields{Shots: shots},
			args:    args{0, 5},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Series{
				XMLName:                              tt.fields.XMLName,
				DsbPunkteS:                           tt.fields.DsbPunkteS,
				TotalScoreOnlyD:                      tt.fields.TotalScoreOnlyD,
				TotalScore:                           tt.fields.TotalScore,
				TotalScoreD:                          tt.fields.TotalScoreD,
				BestDistanceToCenter:                 tt.fields.BestDistanceToCenter,
				OuterShooterTotalScoreBestCount:      tt.fields.OuterShooterTotalScoreBestCount,
				OuterShooterTotalScoreBestCountD:     tt.fields.OuterShooterTotalScoreBestCountD,
				OuterShooterTotalScoreBestCountOnlyD: tt.fields.OuterShooterTotalScoreBestCountOnlyD,
				Shots:                                tt.fields.Shots,
			}
			tt.wantErr(t, s.validateStartAndEnd(tt.args.start, tt.args.end))
		})
	}
}
