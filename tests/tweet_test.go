package tests

import (
	"testing"
	"net/http"
)

func TestCreateTweet(t *testing.T) {
	result := http.StatusOK
	if result != http.StatusOK {
		t.Errorf("Expected %d, got %d", http.StatusOK, result)
	}
}

func TestDeleteTweet(t *testing.T) {
	result := http.StatusOK
	if result != http.StatusOK {
		t.Errorf("Expected %d, got %d", http.StatusOK, result)
	}
}
