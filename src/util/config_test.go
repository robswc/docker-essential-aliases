package util

import "testing"

func TestSaveAndGet(t *testing.T) {

	SaveToTemp("TEST_VAR", "10")

	got := GetFromTemp("TEST_VAR")
	want := "10"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
