package transport

import (
	"reflect"
	"testing"
)

func TestTransport(t *testing.T) {

	type args struct {
		way string
	}

	tests := []struct {
		name string
		args
		want      Transport
		wantError bool
	}{
		{"sea way", args{way: "sea"}, ship{}, false},
		{"road way", args{way: "road"}, truck{}, false},
		{"other way", args{way: "something"}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateTransportWith(tt.args.way)
			if (err != nil) != tt.wantError {
				t.Errorf("CreateTransportWith(%s) error=%v want=%#v got=%#v", tt.args.way, err, tt.want, got)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTransportWith() got = %v, want %v", got, tt.want)
			}
		})
	}

}
