package clothes

import (
	"reflect"
	"testing"
)

func TestGetClothes(t *testing.T) {
	type args struct {
		brand string
	}
	tests := []struct {
		name    string
		args    args
		want    ClothesFactory
		wantErr bool
	}{
		{"adidas clothes", args{brand: "adidas"}, adidasFactory{}, false},
		{"nike clothes", args{brand: "nike"}, nikeFactory{}, false},
		{"vans clothes", args{brand: "vans"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetClothes(tt.args.brand)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClothes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClothes() got = %v, want %v", got, tt.want)
			}
		})
	}
}
