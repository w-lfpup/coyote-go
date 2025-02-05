package coyote

import "testing"

func TesBlep(t *testing.T) {
	var result = Blep("blepp!!")
	if nil == &result {
		t.Error("test failed.")
	}
}
