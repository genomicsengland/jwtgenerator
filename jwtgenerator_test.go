package main

import "testing"

func TestGenerateToken(t *testing.T) {
	expected := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiI1ZXVRb1RzRDNmUURwb0VuWHFvUU40SmlmUFh3MXU1ZyJ9.zDKa5uucLmYUPfmIr7o9_tOgWfgyZyYZgo6k6YQ-mlg"
	token := generateToken("5euQoTsD3fQDpoEnXqoQN4JifPXw1u5g", "m2mU6P2ioGzUonMjwnMRQTMrgY3lI0NV")
	if token != expected {
		t.Errorf("Token was incorrect, got: %s, want: %s.", token, expected)
	}
}
