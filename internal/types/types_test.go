package types

import (
	"testing"
)

func TestEmailIsValid(t *testing.T) {

	tests := []struct {
		email    Email
		expected bool
	}{
		{Email("test@example.com"), true},
		{Email("invalid-email"), false},
		{Email("another@test.co"), true},
		{Email("wrong@.com"), false},
	}

	for _, test := range tests {
		result := test.email.IsValid()
		if result != test.expected {
			t.Errorf("Expected %v for email %s, but got %v", test.expected, test.email, result)
		}
	}
}

func TestGenerateUUID(t *testing.T) {
	uuid1 := GenerateUUID()
	uuid2 := GenerateUUID()
	if uuid1 == uuid2 {
		t.Errorf("Expected different UUIDs, but got %s and %s", uuid1, uuid2)
	}
}

func TestAccountValidate(t *testing.T) {

	tests := []struct {
		account  Account
		expected bool
		message  string
	}{
		{Account{Email: "test@example.com", Password: "password"}, true, ""},
		{Account{Email: "invalid-email", Password: "password"}, false, "Invalid email format"},
		{Account{Email: "test@example.com", Password: "p"}, false, "Small password"},
		{Account{Email: "", Password: "password"}, false, "Invalid email len"},
	}

	for _, test := range tests {
		result, msg := test.account.Validate()
		if result != test.expected || msg != test.message {
			t.Errorf("Expected (%v, %s) for account %v, but got (%v, %s)", test.expected, test.message, test.account, result, msg)
		}
	}
}
