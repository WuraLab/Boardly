package sample

import "testing"

func TestGreetsGitHub(t *testing.T) {
	result := Greet()
	if result != "Testing GHA for Boardly" {
		t.Errorf("Greet() = %s; Testing GHA for Boardly", result)
	}
}