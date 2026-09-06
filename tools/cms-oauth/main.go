package main

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const stateCookie = "cms_oauth_state"

type config struct {
	clientID      string
	clientSecret  string
	publicURL     string
	allowedOrigin string
	listen        string
	scope         string
	allowHTTP     bool
}

type server struct {
	cfg        config
	httpClient *http.Client
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}
	s := &server{cfg: cfg, httpClient: &http.Client{Timeout: 15 * time.Second}}
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", s.health)
	mux.HandleFunc("/auth", s.auth)
	mux.HandleFunc("/callback", s.callback)

	log.Printf("CMS OAuth proxy listening on %s", cfg.listen)
	log.Fatal(http.ListenAndServe(cfg.listen, securityHeaders(mux)))
}

func loadConfig() (config, error) {
	cfg := config{
		clientID:      os.Getenv("GITHUB_CLIENT_ID"),
		clientSecret:  os.Getenv("GITHUB_CLIENT_SECRET"),
		publicURL:     strings.TrimRight(os.Getenv("OAUTH_PUBLIC_URL"), "/"),
		allowedOrigin: strings.TrimRight(os.Getenv("CMS_ALLOWED_ORIGIN"), "/"),
		listen:        getenv("OAUTH_LISTEN", "127.0.0.1:3000"),
		scope:         getenv("GITHUB_SCOPE", "public_repo user:email"),
		allowHTTP:     os.Getenv("ALLOW_INSECURE_HTTP") == "true",
	}
	if cfg.clientID == "" || cfg.clientSecret == "" || cfg.publicURL == "" || cfg.allowedOrigin == "" {
		return cfg, errors.New("GITHUB_CLIENT_ID, GITHUB_CLIENT_SECRET, OAUTH_PUBLIC_URL and CMS_ALLOWED_ORIGIN are required")
	}
	for _, raw := range []string{cfg.publicURL, cfg.allowedOrigin} {
		u, err := url.Parse(raw)
		if err != nil || u.Host == "" {
			return cfg, fmt.Errorf("invalid URL %q", raw)
		}
		if u.Scheme != "https" && !cfg.allowHTTP {
			return cfg, fmt.Errorf("HTTPS is required for %q", raw)
		}
	}
	return cfg, nil
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func (s *server) health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok\n"))
}

func (s *server) auth(w http.ResponseWriter, r *http.Request) {
	state, err := randomState()
	if err != nil {
		http.Error(w, "unable to start authentication", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name: stateCookie, Value: state, Path: "/callback", MaxAge: 600,
		HttpOnly: true, Secure: !s.cfg.allowHTTP, SameSite: http.SameSiteLaxMode,
	})
	callbackURL := s.cfg.publicURL + "/callback"
	params := url.Values{
		"client_id": {s.cfg.clientID}, "redirect_uri": {callbackURL},
		"scope": {s.cfg.scope}, "state": {state}, "allow_signup": {"true"},
	}
	http.Redirect(w, r, "https://github.com/login/oauth/authorize?"+params.Encode(), http.StatusFound)
}

func (s *server) callback(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(stateCookie)
	state := r.URL.Query().Get("state")
	if err != nil || state == "" || subtle.ConstantTimeCompare([]byte(cookie.Value), []byte(state)) != 1 {
		http.Error(w, "invalid OAuth state", http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &http.Cookie{Name: stateCookie, Value: "", Path: "/callback", MaxAge: -1, HttpOnly: true, Secure: !s.cfg.allowHTTP, SameSite: http.SameSiteLaxMode})
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "GitHub did not return an authorization code", http.StatusBadRequest)
		return
	}
	token, err := s.exchangeToken(code)
	if err != nil {
		log.Printf("OAuth token exchange failed: %v", err)
		http.Error(w, "GitHub authentication failed", http.StatusBadGateway)
		return
	}
	s.renderCallback(w, token)
}

func (s *server) exchangeToken(code string) (string, error) {
	payload, _ := json.Marshal(map[string]string{
		"client_id": s.cfg.clientID, "client_secret": s.cfg.clientSecret,
		"code": code, "redirect_uri": s.cfg.publicURL + "/callback",
	})
	req, err := http.NewRequest(http.MethodPost, "https://github.com/login/oauth/access_token", strings.NewReader(string(payload)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "gzu-ai-website-cms-oauth")
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var result struct{ AccessToken, Error, ErrorDescription string }
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK || result.AccessToken == "" {
		return "", fmt.Errorf("GitHub returned %s: %s %s", resp.Status, result.Error, result.ErrorDescription)
	}
	return result.AccessToken, nil
}

var callbackTemplate = template.Must(template.New("callback").Parse(`<!doctype html>
<html><head><meta charset="utf-8"><title>GitHub 登录完成</title></head><body>
<p>正在完成 GitHub 登录……</p>
<script>
(function () {
  const allowedOrigin = {{printf "%q" .Origin}};
  const payload = {{printf "%q" .Payload}};
  function receiveMessage(event) {
    if (event.origin !== allowedOrigin || !window.opener) return;
    window.opener.postMessage(payload, allowedOrigin);
    window.removeEventListener("message", receiveMessage, false);
    window.close();
  }
  window.addEventListener("message", receiveMessage, false);
  if (window.opener) window.opener.postMessage("authorizing:github", allowedOrigin);
})();
</script></body></html>`))

func (s *server) renderCallback(w http.ResponseWriter, token string) {
	content, _ := json.Marshal(map[string]string{"token": token, "provider": "github"})
	data := struct{ Origin, Payload string }{
		Origin:  s.cfg.allowedOrigin,
		Payload: "authorization:github:success:" + string(content),
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	if err := callbackTemplate.Execute(w, data); err != nil {
		log.Printf("render callback: %v", err)
	}
}

func randomState() (string, error) {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(buf), nil
}

func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Referrer-Policy", "no-referrer")
		next.ServeHTTP(w, r)
	})
}
