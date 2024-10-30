package result

import (
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDisagDateTime_parseString(t *testing.T) {
	type args struct {
		str string
	}

	expectedValue := time.Date(2023, 12, 23, 20, 44, 55, 0, time.UTC)

	tests := []struct {
		name    string
		d       DisagDateTime
		args    args
		want    *time.Time
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "Empty string returns zero value",
			d:       DisagDateTime{},
			args:    args{str: ""},
			want:    &time.Time{},
			wantErr: assert.NoError,
		},
		{
			name:    "Malformed input returns an error",
			d:       DisagDateTime{},
			args:    args{str: "2023-12-31"},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name:    "Good case",
			d:       DisagDateTime{},
			args:    args{str: "23.12.2023 20:44:55"},
			want:    &expectedValue,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.parseString(tt.args.str)
			if !tt.wantErr(t, err, fmt.Sprintf("parseString(%v)", tt.args.str)) {
				return
			}
			assert.Equalf(t, tt.want, got, "parseString(%v)", tt.args.str)
		})
	}
}

func TestDisagDateTime_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name      string
		d         DisagDateTime
		args      args
		wantValue DisagDateTime
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name:      "Good case",
			d:         DisagDateTime{},
			args:      args{attr: xml.Attr{Value: "23.12.2023 20:44:55"}},
			wantValue: DisagDateTime(time.Date(2023, 12, 23, 20, 44, 55, 0, time.UTC)),
			wantErr:   assert.NoError,
		},
		{
			name:    "Bad case",
			d:       DisagDateTime{},
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
