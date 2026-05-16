package storage

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/shotomorisaki/kaggle_pacakge/cli/get_discussion/internal/discussion"
)

func intPtr(v int) *int {
	return &v
}

func TestSlugifyTitle(t *testing.T) {
	got := slugifyTitle("Hello, World!!")
	if got != "hello_world" {
		t.Fatalf("unexpected slug: %s", got)
	}
}

func TestFrontMatterRoundTrip(t *testing.T) {
	d := &discussion.Discussion{
		Title:         "My Title",
		Link:          "https://example.com",
		Author:        "Author",
		Comments:      "5",
		PublishedDate: "2024-01-01",
		ContentMD:     "Body",
	}
	content := buildFrontMatter(d) + "Body\n"

	dir := t.TempDir()
	path := filepath.Join(dir, "test.md")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write failed: %v", err)
	}
	meta := readFrontMatter(path)
	if meta["title"] != d.Title {
		t.Fatalf("title mismatch: %s", meta["title"])
	}
	if meta["link"] != d.Link {
		t.Fatalf("link mismatch: %s", meta["link"])
	}
}

func TestEnsureUniquePath(t *testing.T) {
	dir := t.TempDir()
	base := "discussion"
	first := filepath.Join(dir, base+".md")
	if err := os.WriteFile(first, []byte("x"), 0o644); err != nil {
		t.Fatalf("write failed: %v", err)
	}

	paths := map[string]struct{}{first: {}}
	got := ensureUniquePath(dir, base, paths)
	if got == first {
		t.Fatalf("expected unique path, got %s", got)
	}
}

func TestLoadIgnoreRules(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "ignore.yml")
	content := `- postId: 123
  author: "Alice"
  link: "https://www.kaggle.com/discussion/123"
- author: "Bob"
`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write failed: %v", err)
	}

	rules, err := LoadIgnoreRules(path)
	if err != nil {
		t.Fatalf("LoadIgnoreRules failed: %v", err)
	}
	if len(rules) != 2 {
		t.Fatalf("expected 2 rules, got %d", len(rules))
	}
	if rules[0].PostID == nil || *rules[0].PostID != 123 {
		t.Fatalf("unexpected post id: %+v", rules[0].PostID)
	}
	if rules[0].Author != "Alice" {
		t.Fatalf("unexpected author: %q", rules[0].Author)
	}
}

func TestSaveDiscussionHonorsIgnoreRules(t *testing.T) {
	dir := t.TempDir()
	d := &discussion.Discussion{
		PostID:    123,
		Title:     "Ignored Title",
		Link:      "https://www.kaggle.com/discussion/123",
		Author:    "Alice",
		ContentMD: "Body",
	}

	path, err := SaveDiscussion(d, dir, map[string]string{}, []IgnoreRule{{PostID: intPtr(123)}})
	if !errors.Is(err, ErrDiscussionIgnored) {
		t.Fatalf("expected ErrDiscussionIgnored, got %v", err)
	}
	if path != "" {
		t.Fatalf("expected empty path, got %q", path)
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("ReadDir failed: %v", err)
	}
	if len(entries) != 0 {
		t.Fatalf("expected no files to be written, got %d entries", len(entries))
	}
}

func TestIgnoreRuleMatchingByAuthorAndLink(t *testing.T) {
	d := &discussion.Discussion{
		PostID:    456,
		Title:     "Kept Title",
		Link:      "https://www.kaggle.com/discussion/456?foo=bar#baz",
		Author:    "Alice",
		ContentMD: "Body",
	}

	rules := []IgnoreRule{
		{Author: "Bob"},
		{Link: "https://www.kaggle.com/discussion/456"},
	}

	if !shouldIgnoreDiscussion(d, rules) {
		t.Fatalf("expected discussion to be ignored by link rule")
	}
}
