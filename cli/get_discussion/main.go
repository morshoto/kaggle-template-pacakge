package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/shotomorisaki/kaggle_pacakge/cli/get_discussion/internal/api"
	"github.com/shotomorisaki/kaggle_pacakge/cli/get_discussion/internal/client"
	"github.com/shotomorisaki/kaggle_pacakge/cli/get_discussion/internal/discussion"
	"github.com/shotomorisaki/kaggle_pacakge/cli/get_discussion/internal/storage"
	"github.com/shotomorisaki/kaggle_pacakge/cli/get_discussion/pkg/urlutil"
)

func fetchCandidateLimit(limit int, all bool, ignoreRules []storage.IgnoreRule) int {
	if all {
		return 0
	}
	if len(ignoreRules) > 0 && limit > 0 {
		return 0
	}
	return limit
}

func main() {
	var (
		link       string
		sort       string
		timeFilter string
		outputDir  string
		delay      float64
		verbose    bool
		limit      int
		all        bool
	)

	flag.StringVar(&link, "link", "", "Download a single discussion by URL.")
	flag.StringVar(&sort, "sort", "hotness", "Sort: hotness, recent_comments, recently_posted, most_votes, most_comments.")
	flag.StringVar(&timeFilter, "time-filter", "", "Time filter: last_30_days, last_7_days, today.")
	flag.StringVar(&outputDir, "output-dir", "discussion", "Output directory for Markdown files.")
	flag.Float64Var(&delay, "delay", 0.5, "Delay in seconds between requests.")
	flag.IntVar(&limit, "limit", 10, "Max discussions to download (default 10).")
	flag.BoolVar(&all, "all", false, "Download all discussions (ignores --limit).")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose logging.")
	flag.Parse()

	storage.LoadEnvFile(".env")

	ignoreRules, err := storage.LoadIgnoreRules(
		filepath.Join("cli", "get_discussion", "ignore.yml"),
		"ignore.yml",
	)
	if err != nil {
		log.Printf("[warn] Failed to load ignore rules: %v", err)
	}

	httpClient := client.NewClient(verbose)
	fetchLimit := fetchCandidateLimit(limit, all, ignoreRules)

	var urls []string

	if link != "" {
		urls = []string{urlutil.CanonicalizeURL(link)}
	} else {
		sortKey := urlutil.NormalizeChoice(sort)
		timeKey := urlutil.NormalizeChoice(timeFilter)

		if sortKey != "" {
			if _, ok := urlutil.SortParam(sortKey); !ok {
				log.Fatalf("Unknown sort option: %s", sort)
			}
		}
		if timeKey != "" {
			if _, ok := urlutil.TimeFilterParam(timeKey); !ok {
				log.Fatalf("Unknown time filter: %s", timeFilter)
			}
		}

		competition := os.Getenv("COMPETITION")

		if competition != "" {
			forumID, err := api.FetchCompetitionForumID(httpClient, competition)
			if err != nil {
				log.Printf("[warn] Competition API failed: %v", err)
			} else {
				urls, err = api.FetchTopicListByForumID(httpClient, forumID, sortKey, timeKey, fetchLimit)
				if err != nil {
					log.Printf("[warn] Topic list API failed: %v", err)
					urls = nil
				}
				if len(urls) == 0 {
					log.Printf("[warn] No topics found via forumId=%d competition=%s", forumID, competition)
				}
			}
		}

		if len(urls) == 0 {
			listingURL := urlutil.BuildListingURL(sortKey, timeKey)
			body, err := httpClient.FetchBody(listingURL, nil)
			if err != nil {
				log.Printf("[warn] Failed to fetch listing: %v", err)
			} else {
				urls = discussion.ExtractDiscussionLinksFromHTML(body, listingURL)
				if fetchLimit > 0 && len(urls) > fetchLimit {
					urls = urls[:fetchLimit]
				}
			}

			if len(urls) == 0 && competition != "" {
				compURL := urlutil.BuildCompetitionListingURL(competition, sortKey, timeKey)
				httpClient.LogInfo("Retrying with competition listing url=%s", compURL)
				body, err = httpClient.FetchBody(compURL, nil)
				if err != nil {
					log.Printf("[warn] Competition listing failed: %v", err)
				} else {
					urls = discussion.ExtractDiscussionLinksFromHTML(body, compURL)
					if fetchLimit > 0 && len(urls) > fetchLimit {
						urls = urls[:fetchLimit]
					}
					if len(urls) == 0 {
						log.Printf("[warn] No discussion links found on competition listing url=%s", compURL)
					}
				}
			}
		}
	}

	existingByLink := storage.LoadExistingLinks(outputDir)
	d := time.Duration(float64(time.Second) * delay)
	savedCount := 0

	for discussionItem := range discussion.IterDiscussions(urls, httpClient, d) {
		path, err := storage.SaveDiscussion(discussionItem, outputDir, existingByLink, ignoreRules)
		if err != nil {
			if errors.Is(err, storage.ErrDiscussionIgnored) {
				continue
			}
			log.Printf("[warn] Failed to save %s: %v", discussionItem.Link, err)
			continue
		}
		fmt.Println(path)
		savedCount++
		if !all && limit > 0 && savedCount >= limit {
			break
		}
	}
}
