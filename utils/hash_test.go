package utils

import "testing"

func TestHash(t *testing.T) {
	password := "password"
	hashed_password := Hash(password)
	if !CheckHash(hashed_password, password) {
		t.Fatalf("Password failed to match")
	}
}
