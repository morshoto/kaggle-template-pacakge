package urlutil

import (
	"strings"
	"testing"
)

func TestEnsureURL(t *testing.T) {
	got := EnsureURL("kaggle.com/discussion/123")
	if got != "https://kaggle.com/discussion/123" {
		t.Fatalf("unexpected url: %s", got)
	}
}

func TestCanonicalizeURL(t *testing.T) {
	got := CanonicalizeURL("https://kaggle.com/discussion/123?foo=bar#section")
	if got != "https://kaggle.com/discussion/123" {
		t.Fatalf("unexpected canonical: %s", got)
	}
}

func TestExtractTopicID(t *testing.T) {
	id, ok := ExtractTopicID("https://www.kaggle.com/discussion/98765/foo")
	if !ok || id != 98765 {
		t.Fatalf("unexpected id: %v ok=%v", id, ok)
	}
}

func TestBuildListingURL(t *testing.T) {
	got := BuildListingURL("hotness", "last_7_days")
	if got == "https://www.kaggle.com/discussions" {
		t.Fatalf("expected params in listing url")
	}
}

func TestBuildCompetitionListingURL(t *testing.T) {
	got := BuildCompetitionListingURL("titanic", "most_votes", "")
	if got == "" {
		t.Fatalf("expected competition url")
	}
	if !strings.HasPrefix(got, "https://www.kaggle.com/competitions/titanic/discussion") {
		t.Fatalf("unexpected competition url: %s", got)
	}
}

func TestNormalizeChoice(t *testing.T) {
	if got := NormalizeChoice("Recent Comments"); got != "recent_comments" {
		t.Fatalf("unexpected normalized: %s", got)
	}
}
