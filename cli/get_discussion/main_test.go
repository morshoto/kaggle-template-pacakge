package main

import (
	"testing"

	"github.com/shotomorisaki/kaggle_pacakge/cli/get_discussion/internal/storage"
)

func TestFetchCandidateLimit(t *testing.T) {
	if got := fetchCandidateLimit(45, false, nil); got != 45 {
		t.Fatalf("expected 45, got %d", got)
	}

	if got := fetchCandidateLimit(45, true, nil); got != 0 {
		t.Fatalf("expected 0 for all mode, got %d", got)
	}

	if got := fetchCandidateLimit(45, false, []storage.IgnoreRule{{Author: "Alice"}}); got != 0 {
		t.Fatalf("expected 0 when ignore rules are present, got %d", got)
	}
}
