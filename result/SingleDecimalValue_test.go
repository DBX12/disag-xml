package result

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleDecimalValue_String(t *testing.T) {
	tests := []struct {
		name string
		d    SingleDecimalValue
		want string
	}{
		{
			name: "Zero",
			d:    SingleDecimalValue(0.00),
			want: "0.0",
		},
		{
			name: "Ten point zero",
			d:    SingleDecimalValue(10.0),
			want: "10.0",
		},
		{
			name: "Ten point one two",
			d:    SingleDecimalValue(10.12),
			want: "10.1",
		},
		{
			name: "Rounding does not happen",
			d:    SingleDecimalValue(2.89),
			want: "2.8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.d.String())
		})
	}
}

func TestSingleDecimalValue_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		d       SingleDecimalValue
		args    args
		want    SingleDecimalValue
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "Empty string",
			d:       SingleDecimalValue(-1),
			args:    args{xml.Attr{Value: ""}},
			want:    SingleDecimalValue(-1),
			wantErr: assert.Error,
		},
		{
			name:    "Separated with point",
			d:       SingleDecimalValue(-1),
			args:    args{xml.Attr{Value: "1.2"}},
			want:    SingleDecimalValue(1.2),
			wantErr: assert.NoError,
		},

		{
			name:    "Separated with comma",
			d:       SingleDecimalValue(-1),
			args:    args{xml.Attr{Value: "1,2"}},
			want:    SingleDecimalValue(1.2),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.wantErr(t, tt.d.UnmarshalXMLAttr(tt.args.attr)) {
				return
			}
			assert.Equal(t, tt.want, tt.d)
		})
	}
}
