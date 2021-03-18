package single

import (
	"testing"
)

func Test_newInstance(t *testing.T) {
	for i := 0; i < 30; i++ {
		go newInstance()
	}
}
