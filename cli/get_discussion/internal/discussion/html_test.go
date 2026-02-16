package discussion

import "testing"

func TestExtractDiscussionLinksFromHTML(t *testing.T) {
	html := []byte(`<a href="/discussion/123/test">A</a><a href="/discussions/456">B</a>`)
	links := ExtractDiscussionLinksFromHTML(html, "https://www.kaggle.com/discussions")
	if len(links) != 2 {
		t.Fatalf("expected 2 links, got %d", len(links))
	}
}

func TestExtractTitleFromHTML(t *testing.T) {
	html := []byte(`<html><head><meta property="og:title" content="Meta Title"></head><body><h1>Header Title</h1></body></html>`)
	got := extractTitleFromHTML(html)
	if got != "Header Title" {
		t.Fatalf("unexpected title: %s", got)
	}
}

func TestHTMLToMarkdown(t *testing.T) {
	html := []byte(`<h1>Title</h1><p>Hello <strong>World</strong></p><a href="https://example.com">Link</a>`)
	md := htmlToMarkdown(html)
	if md == "" {
		t.Fatalf("expected markdown content")
	}
}
