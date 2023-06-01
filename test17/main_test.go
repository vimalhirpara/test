package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Tell the client that the API version is 1.3
	w.Header().Add("API-VERSION", "1.3")
	w.Write([]byte("ok"))
}

func TestHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "https://contis.free.beeceptor.com/VirtualAccountV2", nil)
	w := httptest.NewRecorder()

	Handler(w, req)

	// We should get a good status code
	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}

	// Make sure that the version was 1.3
	if want, got := "1.3", w.Result().Header.Get("API-VERSION"); want != got {
		t.Fatalf("expected API-VERSION to be %s, instead got: %s", want, got)
	}
}
