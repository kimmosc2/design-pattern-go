package phone

import "testing"

func TestTypeCPortAdapter(t *testing.T) {
	type args struct {
		p typeCPort
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"android phone",args{p: typeC{}},"micro-usb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &adapter{}
			if got := a.insertTypeCPort(tt.args.p); got != tt.want {
				t.Errorf("insertTypeCPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
