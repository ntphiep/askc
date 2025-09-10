package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendToSlack_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	defer ts.Close()

	err := SendToSlack(ts.URL, "Test message")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestSendToSlack_Failure(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
	}))
	defer ts.Close()

	err := SendToSlack(ts.URL, "Test message")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
