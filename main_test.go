package main

import "testing"

func TestRun(t *testing.T) {
	expected := "setup CI/CD with travis ci "
	got := run()
	if expected != got {
		t.Fatalf("expected %v got %v", expected, got)
	}
}
