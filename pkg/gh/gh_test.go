package gh

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	_, err := NewClient("https://github.enterprise.com/", "token", "bot-app", "bot-app@mail.com")
	if err != nil {
		t.Fatalf("NewClient : %s", err)
	}
}
