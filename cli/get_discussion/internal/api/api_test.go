package api

import (
	"encoding/json"
	"testing"
)

func TestTopicResponseUnmarshal(t *testing.T) {
	payload := []byte(`{"forumTopic":{"name":"Title","url":"/discussion/1","authorUserDisplayName":"User","totalMessages":3,"postDate":"2024-01-01","firstMessageId":10}}`)
	var resp TopicResponse
	if err := json.Unmarshal(payload, &resp); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if resp.ForumTopic.Name != "Title" || resp.ForumTopic.FirstMessageID != 10 {
		t.Fatalf("unexpected data: %+v", resp.ForumTopic)
	}
}

func TestTopicListResponseUnmarshal(t *testing.T) {
	payload := []byte(`{"count":2,"topics":[{"topicUrl":"/discussion/1"},{"url":"/discussion/2"}]}`)
	var resp TopicListResponse
	if err := json.Unmarshal(payload, &resp); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if resp.Count != 2 || len(resp.Topics) != 2 {
		t.Fatalf("unexpected count: %+v", resp)
	}
}
