package main

import "testing"

func TestGenerateToken(t *testing.T) {
	expected := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiI1ZXVRb1RzRDNmUURwb0VuWHFvUU40SmlmUFh3MXU1ZyIsICJqdGkiOiJlNWJhNzU5MDE1NmUzMzNlZjlhYTRiOTYxNmE1NTkyMSIsICJvaWQiOiJiaW8ifQ.Ar7ssnoN9uGVRYuDMwBxIZaNV1tZI1caYhUrsD9OWtc"
	token := generateToken("5euQoTsD3fQDpoEnXqoQN4JifPXw1u5g", "m2mU6P2ioGzUonMjwnMRQTMrgY3lI0NV", "bio")
	if token != expected {
		t.Errorf("Token was incorrect, got: %s, want: %s.", token, expected)
	}
}
