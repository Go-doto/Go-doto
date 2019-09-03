package dota_api

import (
	"testing"
)

func TestNewClientWithToken(t *testing.T) {
	_, err := NewClientWithToken("")
	if err == nil {
		t.Error("expected error but nil given when token is empty")
	}

	_, err = NewClientWithToken("123")
	if err != nil {
		t.Error("unexpected error")
	}
}
