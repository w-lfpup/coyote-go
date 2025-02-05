package coyote

import "testing"

func TestGetKnownResult(t *testing.T) {
	var result = Blep("hello!")
	if nil == &result {
		t.Error("test failed.")
	}
}
