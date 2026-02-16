package client

import (
	"net/http"
	"net/url"
	"testing"
)

func TestSimpleCookieJar(t *testing.T) {
	jar := &simpleCookieJar{cookies: map[string]string{}}
	u, _ := url.Parse("https://example.com")
	jar.SetCookies(u, []*http.Cookie{{Name: "XSRF-TOKEN", Value: "abc"}})
	cookies := jar.Cookies(u)
	if len(cookies) != 1 {
		t.Fatalf("expected 1 cookie, got %d", len(cookies))
	}
	if cookies[0].Name != "XSRF-TOKEN" || cookies[0].Value != "abc" {
		t.Fatalf("unexpected cookie: %+v", cookies[0])
	}
}
