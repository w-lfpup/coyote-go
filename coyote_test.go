package coyote

import "testing"

func TestGetKnownResult(t *testing.T) {
    Log("hello!");
    if false != true {
        t.Error("Known result not received, test failed.")
    }
}
