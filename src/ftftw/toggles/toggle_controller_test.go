package toggles

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	getToggle("I am a passing test")
}

func TestFail(t *testing.T) {
	t.Errorf("%s", "I am a failing test")
}
