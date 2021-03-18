package gui

import (
	"reflect"
	"testing"
)

func TestNewUIFactory(t *testing.T) {
	type args struct {
		platform string
	}
	tests := []struct {
		name    string
		args    args
		want    UIFactory
		wantErr bool
	}{
		{"windows", args{platform: "win"}, winUI{}, false},
		{"web", args{platform: "web"}, webUI{}, false},
		{"linux", args{platform: "linux"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUIFactory(tt.args.platform)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUIFactory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUIFactory() got = %v, want %v", got, tt.want)
			}
		})
	}
}
