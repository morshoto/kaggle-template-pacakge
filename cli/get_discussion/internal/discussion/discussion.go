package discussion

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/shotomorisaki/kaggle_pacakge/cli/get_discussion/internal/api"
	"github.com/shotomorisaki/kaggle_pacakge/cli/get_discussion/internal/client"
	"github.com/shotomorisaki/kaggle_pacakge/cli/get_discussion/pkg/urlutil"
)

// Discussion holds all metadata and content for a single Kaggle discussion.
type Discussion struct {
	Title         string
	Link          string
	Author        string
	Comments      string
	PublishedDate string
	ContentMD     string
}

func BuildDiscussionFromAPI(c *client.Client, rawURL string, topicID int) (*Discussion, error) {
	// Warm up cookies.
	_, _ = c.FetchBody(rawURL, nil)

	topicResp, err := api.FetchTopicData(c, topicID)
	if err != nil {
		return nil, err
	}
	t := topicResp.ForumTopic
	if t.Name == "" {
		return nil, fmt.Errorf("empty forumTopic for topic_id=%d", topicID)
	}

	msgResp, err := api.FetchTopicMessages(c, topicID)
	if err != nil {
		return nil, err
	}

	contentMD := buildDiscussionMarkdown(msgResp, t.FirstMessageID)

	link := t.URL
	if link == "" {
		link = rawURL
	}
	if !strings.HasPrefix(link, "http") {
		base, _ := url.Parse("https://www.kaggle.com")
		ref, _ := url.Parse(link)
		link = base.ResolveReference(ref).String()
	}
	link = urlutil.CanonicalizeURL(link)

	author := t.AuthorUserDisplayName
	if author == "" {
		author = t.AuthorUserName
	}

	comments := ""
	if t.TotalMessages != nil {
		comments = fmt.Sprint(*t.TotalMessages)
	}

	return &Discussion{
		Title:         urlutil.FirstNonEmpty(t.Name, "untitled_discussion"),
		Link:          link,
		Author:        author,
		Comments:      comments,
		PublishedDate: t.PostDate,
		ContentMD:     strings.TrimSpace(contentMD),
	}, nil
}

func buildDiscussionMarkdown(msgResp *api.MessagesResponse, firstMessageID int) string {
	if msgResp == nil || len(msgResp.Comments) == 0 {
		return ""
	}

	var mainMessage string
	var replies []string

	for i, m := range msgResp.Comments {
		body := strings.TrimSpace(urlutil.FirstNonEmpty(m.RawMarkdown, m.Content))
		if body == "" {
			continue
		}

		if m.ID == firstMessageID || (firstMessageID == 0 && i == 0) {
			if mainMessage == "" {
				mainMessage = body
				continue
			}
		}

		author := commentAuthor(m)
		if author == "" {
			author = "Unknown"
		}
		replies = append(replies, fmt.Sprintf("## Comment by %s\n\n%s", author, body))
	}

	if mainMessage == "" {
		mainMessage = strings.TrimSpace(urlutil.FirstNonEmpty(
			msgResp.Comments[0].RawMarkdown,
			msgResp.Comments[0].Content,
		))
		if len(replies) > 0 {
			replies = replies[1:]
		}
	}
	if mainMessage == "" {
		return ""
	}
	if len(replies) == 0 {
		return mainMessage
	}
	return mainMessage + "\n\n---\n\n" + strings.Join(replies, "\n\n---\n\n")
}

func commentAuthor(m api.ForumComment) string {
	return urlutil.FirstNonEmpty(
		m.AuthorUserDisplayName,
		m.AuthorUserName,
		m.User.DisplayName,
		m.User.UserName,
		m.User.Name,
	)
}

// IterDiscussions yields Discussion values for each URL, with API -> HTML fallback.
func IterDiscussions(urls []string, c *client.Client, delay time.Duration) <-chan *Discussion {
	ch := make(chan *Discussion)
	go func() {
		defer close(ch)
		for _, rawURL := range urls {
			topicID, hasID := urlutil.ExtractTopicID(rawURL)
			var d *Discussion
			var err error

			if hasID {
				d, err = BuildDiscussionFromAPI(c, rawURL, topicID)
				if err != nil {
					log.Printf("[warn] API failed for %s: %v — falling back to HTML", rawURL, err)
					d, err = BuildDiscussionFromHTML(c, rawURL)
				}
			} else {
				log.Printf("[warn] No topic ID detected in URL %s — using HTML parser", rawURL)
				d, err = BuildDiscussionFromHTML(c, rawURL)
			}

			if err != nil {
				log.Printf("[warn] Skipping %s: %v", rawURL, err)
				continue
			}
			if strings.TrimSpace(d.ContentMD) == "" {
				log.Printf("[warn] Empty content for %s", rawURL)
			}
			ch <- d
			if delay > 0 {
				time.Sleep(delay)
			}
		}
	}()
	return ch
}
