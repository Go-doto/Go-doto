package dota_api

import (
	"errors"
	"testing"
)

func TestNewClientWithToken(t *testing.T) {
	_, err := NewClientWithToken("")
	if err == nil {
		t.Error("expected error but nil given when token is empty")
	} else if !errors.As(err, &ValidationError{}) {
		t.Error("expected validation error")
	}
	_, err = NewClientWithToken("123")
	if err != nil {
		t.Error("unexpected error")
	}
}
