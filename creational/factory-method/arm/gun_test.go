package arm

import (
	"reflect"
	"testing"
)

func TestNewGun(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    Arm
		wantErr bool
	}{
		{"ak47", args{name: "ak47"}, &ak47{}, false},
		{"m4a1", args{name: "m4a1"}, &m4a1{}, false},
		{"awm", args{name: "awm"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGun(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGun() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGun() got = %v, want %v", got, tt.want)
			}
		})
	}
}
