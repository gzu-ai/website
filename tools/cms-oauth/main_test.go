package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func testServer() *server {
	return &server{cfg: config{
		clientID:      "client-id",
		clientSecret:  "client-secret",
		publicURL:     "https://lab.example.edu",
		allowedOrigin: "https://lab.example.edu",
		scope:         "public_repo user:email",
	}}
}

func TestRandomState(t *testing.T) {
	first, err := randomState()
	if err != nil {
		t.Fatal(err)
	}
	second, err := randomState()
	if err != nil {
		t.Fatal(err)
	}
	if len(first) < 40 || first == second {
		t.Fatalf("states are not sufficiently random: %q %q", first, second)
	}
}

func TestAuthRedirectAndCookie(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "https://lab.example.edu/auth", nil)
	recorder := httptest.NewRecorder()
	testServer().auth(recorder, req)

	if recorder.Code != http.StatusFound {
		t.Fatalf("status = %d", recorder.Code)
	}
	location, err := url.Parse(recorder.Header().Get("Location"))
	if err != nil {
		t.Fatal(err)
	}
	if location.Host != "github.com" || location.Query().Get("client_id") != "client-id" {
		t.Fatalf("unexpected redirect: %s", location)
	}
	if location.Query().Get("redirect_uri") != "https://lab.example.edu/callback" {
		t.Fatalf("unexpected callback: %s", location.Query().Get("redirect_uri"))
	}
	cookies := recorder.Result().Cookies()
	if len(cookies) != 1 || cookies[0].Name != stateCookie || !cookies[0].Secure || !cookies[0].HttpOnly {
		t.Fatalf("unexpected state cookie: %#v", cookies)
	}
}

func TestCallbackRejectsInvalidState(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "https://lab.example.edu/callback?state=wrong&code=x", nil)
	req.AddCookie(&http.Cookie{Name: stateCookie, Value: "expected"})
	recorder := httptest.NewRecorder()
	testServer().callback(recorder, req)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("status = %d", recorder.Code)
	}
}

func TestCallbackMessageUsesOnlyAllowedOrigin(t *testing.T) {
	recorder := httptest.NewRecorder()
	testServer().renderCallback(recorder, "secret-token")
	body := recorder.Body.String()
	if !strings.Contains(body, "authorization:github:success") || !strings.Contains(body, "lab.example.edu") {
		t.Fatalf("callback response missing expected values: %s", body)
	}
	if strings.Contains(body, "postMessage(payload, \"*\")") {
		t.Fatal("callback must not use a wildcard origin")
	}
}
