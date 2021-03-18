package house

import (
	"reflect"
	"testing"
)

func TestDirector(t *testing.T) {
	var tests = []struct {
		name      string
		builder   iBuilder
		wantHouse house
		wantError bool
	}{
		{"wooden builder", &woodenBuilder{}, house{floor: 1, windowType: "wood", doorType: "wood", wallType: "wood"}, false},
		{"steel builder", &steelBuilder{}, house{floor: 1, windowType: "steel", doorType: "steel", wallType: "steel"}, false},
		{"diamond builder", nil, house{}, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			director := newDirector(test.builder)
			h, err := director.buildHouse()
			if (err != nil) != test.wantError {
				t.Errorf("director builder: wantError = %v, error = %v", test.wantError, err)
				return
			}
			if !reflect.DeepEqual(h, test.wantHouse) {
				t.Errorf("director builder: %T, want = %#v , got = %#v", test.builder, test.wantHouse, h)
			}
		})
	}
}
