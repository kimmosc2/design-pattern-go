package once

import (
	"testing"
)

func TestOnce(t *testing.T) {
	for i := 0; i < 30; i++ {
		go newInstance()
	}
}
