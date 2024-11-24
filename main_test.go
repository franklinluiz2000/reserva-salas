package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIsAdmin(t *testing.T) {
	users = []User{
		{Username: "admin", Password: "1234", Admin: true},
		{Username: "user", Password: "abcd", Admin: false},
	}

	tests := []struct {
		username string
		expected bool
	}{
		{"admin", true},
		{"user", false},
		{"unknown", false},
	}

	for _, test := range tests {
		result := isAdmin(test.username)
		if result != test.expected {
			t.Errorf("isAdmin(%q) = %v; want %v", test.username, result, test.expected)
		}
	}
}

func TestHandleLogin(t *testing.T) {
	req := httptest.NewRequest("POST", "/login", nil)
	req.Form = map[string][]string{
		"username": {"admin"},
		"password": {"1234"},
	}

	users = []User{
		{Username: "admin", Password: "1234", Admin: true},
	}

	rr := httptest.NewRecorder()
	handleLogin(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("handleLogin() = %v; want %v", rr.Code, http.StatusSeeOther)
	}

	location := rr.Header().Get("Location")
	if location != "/menu" {
		t.Errorf("Redirected to %q; want /menu", location)
	}
}
