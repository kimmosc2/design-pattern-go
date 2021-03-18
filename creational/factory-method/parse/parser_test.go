package parse

import (
	"reflect"
	"testing"
)

func TestNewParser(t *testing.T) {
	type args struct {
		parserType string
	}
	tests := []struct {
		name    string
		args    args
		want    Parser
		wantErr bool
	}{
		{"json", args{"json"}, json{}, false},
		{"yaml", args{"yaml"}, yaml{}, false},
		{"toml", args{"toml"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewParser(tt.args.parserType)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewParser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
