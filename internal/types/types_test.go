package types

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
	"time"
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

func normalizeJSON(input string) (string, error) {
	var jsonObj map[string]interface{}
	err := json.Unmarshal([]byte(input), &jsonObj)
	if err != nil {
		return "", err
	}
	normalized, err := json.Marshal(jsonObj)
	if err != nil {
		return "", err
	}
	return string(normalized), nil
}

func TestTokenVerify(t *testing.T) {
	secretKey := "mysecretkey"
	os.Setenv("JWT_SECRET_KEY", secretKey)

	username := "testuser"
	duration := time.Minute

	token := JWT(username, duration)

	valid, err := token.Verify()
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if !valid {
		t.Fatal("Expected token to be valid")
	}

	// Test expired token
	expiredToken := JWT(username, -time.Minute)
	valid, err = expiredToken.Verify()
	if err == nil {
		t.Fatal("Expected error for expired token, but got nil")
	}

	if valid {
		t.Fatal("Expected token to be invalid")
	}

	// Test token with invalid signature
	parts := strings.Split(token.JWT, ".")
	parts[2] = "invalidsignature"
	invalidToken := Token{JWT: strings.Join(parts, ".")}
	valid, err = invalidToken.Verify()
	if err == nil {
		t.Fatal("Expected error for token with invalid signature, but got nil")
	}

	if valid {
		t.Fatal("Expected token to be invalid")
	}
}
