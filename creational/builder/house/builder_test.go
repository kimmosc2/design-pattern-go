package house

import (
	"reflect"
	"testing"
)

func Test_getBuilder(t *testing.T) {
	type args struct {
		texture string
	}
	tests := []struct {
		name    string
		args    args
		want    iBuilder
		wantErr bool
	}{
		{"wood house", args{texture: "wood"}, &woodenBuilder{}, false},
		{"steel house", args{texture: "steel"}, &steelBuilder{}, false},
		{"diamond house", args{texture: "diamond"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getBuilder(tt.args.texture)
			if (err != nil) != tt.wantErr {
				t.Errorf("getBuilder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBuilder() got = %v, want %v", got, tt.want)
			}
		})
	}
}
