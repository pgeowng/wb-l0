package model

import (
	"os"
	"testing"
)

func expectErr(t *testing.T, err error, message string) {
	if err == nil {
		t.Logf("expected err, got nil: %s", message)
		t.Fail()
	}
}

func expectValid(t *testing.T, err error, message string) {
	if err != nil {
		t.Logf("expected valid, got err(%s) for %s", err, message)
		t.Fail()
	}
}

func TestValidate(t *testing.T) {
	bytes, err := os.ReadFile("model.json")
	if err != nil {
		t.Log("bad model.json")
		t.Fail()
		return
	}

	o := Order{}

	err = o.FromJSONBuffer(bytes)
	if err != nil {
		t.Logf("bad json: %s\n", err)
		t.Fail()
		return
	}

	expectValid(t, o.Validate(), "empty")
}
