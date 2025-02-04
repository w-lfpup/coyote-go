package coyote_go

import "testing"

func TestGetKnownResult(t *testing.T) {
    Log("hello!");
    if false != true {
        t.Error("test failed.")
    }
}
