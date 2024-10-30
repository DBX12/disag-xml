package result

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDisagDateIso_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name      string
		d         DisagDateIso
		args      args
		wantValue DisagDateIso
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name:      "Good case",
			d:         DisagDateIso{},
			args:      args{attr: xml.Attr{Value: "20231231"}},
			wantValue: DisagDateIso(time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)),
			wantErr:   assert.NoError,
		},
		{
			name:    "Bad case",
			d:       DisagDateIso{},
			args:    args{attr: xml.Attr{Value: "2023-12-31"}},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.wantErr(t, tt.d.UnmarshalXMLAttr(tt.args.attr)) {
				return
			}
			assert.Equal(t, tt.wantValue, tt.d)
		})
	}
}

func TestDisagDateIso_parseString(t *testing.T) {
	type args struct {
		str string
	}
	expected := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name               string
		d                  DisagDateIso
		args               args
		wantErr            assert.ErrorAssertionFunc
		wantValue          *time.Time
		wantValueAssertion func(t *testing.T, got *time.Time)
	}{
		{
			name:      "Empty string returns time.Zero",
			d:         DisagDateIso{},
			args:      args{""},
			wantErr:   assert.NoError,
			wantValue: &time.Time{},
		},
		{
			name:      "Malformed strings cause an error",
			d:         DisagDateIso{},
			args:      args{"2023-12-31"},
			wantErr:   assert.Error,
			wantValue: nil,
		},
		{
			name:      "Valid strings are parsed",
			d:         DisagDateIso{},
			args:      args{"20231231"},
			wantErr:   assert.NoError,
			wantValue: &expected,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.parseString(tt.args.str)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.wantValue, got)
		})
	}
}
