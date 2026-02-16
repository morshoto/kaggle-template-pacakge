package urlutil

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

const (
	baseListingURL     = "https://www.kaggle.com/discussions"
	competitionListURL = "https://www.kaggle.com/competitions/%s/discussion"
)

var sortOptions = map[string]string{
	"hotness":         "hotness",
	"recent_comments": "recent-comments",
	"recently_posted": "recently-posted",
	"most_votes":      "most-votes",
	"most_comments":   "most-comments",
}

var timeFilterOptions = map[string]string{
	"last_30_days": "last-30-days",
	"last_7_days":  "last-7-days",
	"today":        "today",
}

func EnsureURL(raw string) string {
	if !strings.HasPrefix(raw, "http://") && !strings.HasPrefix(raw, "https://") {
		return "https://" + strings.TrimLeft(raw, "/")
	}
	return raw
}

func CanonicalizeURL(raw string) string {
	u, err := url.Parse(EnsureURL(raw))
	if err != nil {
		return raw
	}
	u.RawQuery = ""
	u.Fragment = ""
	return u.String()
}

var topicIDRegex = regexp.MustCompile(`/discussion/(\d+)`)

func ExtractTopicID(rawURL string) (int, bool) {
	m := topicIDRegex.FindStringSubmatch(rawURL)
	if m == nil {
		return 0, false
	}
	var id int
	fmt.Sscanf(m[1], "%d", &id)
	return id, true
}

func SortParam(key string) (string, bool) {
	v, ok := sortOptions[key]
	return v, ok
}

func TimeFilterParam(key string) (string, bool) {
	v, ok := timeFilterOptions[key]
	return v, ok
}

func BuildListingURL(sortKey, timeKey string) string {
	params := url.Values{}
	if v, ok := SortParam(sortKey); ok {
		params.Set("sort", v)
	}
	if v, ok := TimeFilterParam(timeKey); ok {
		params.Set("time", v)
	}
	if len(params) == 0 {
		return baseListingURL
	}
	return baseListingURL + "?" + params.Encode()
}

func BuildCompetitionListingURL(competition, sortKey, timeKey string) string {
	base := fmt.Sprintf(competitionListURL, strings.TrimSpace(competition))
	params := url.Values{}
	if v, ok := SortParam(sortKey); ok {
		params.Set("sort", v)
	}
	if v, ok := TimeFilterParam(timeKey); ok {
		params.Set("time", v)
	}
	if len(params) == 0 {
		return base
	}
	return base + "?" + params.Encode()
}

func NormalizeChoice(s string) string {
	s = strings.TrimSpace(strings.ToLower(s))
	s = strings.ReplaceAll(s, " ", "_")
	s = strings.ReplaceAll(s, "-", "_")
	return s
}

func FirstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}
