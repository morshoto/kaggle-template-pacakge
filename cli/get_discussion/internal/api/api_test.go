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

func TestMessagesResponseUnmarshal(t *testing.T) {
	payload := []byte(`{"comments":[{"id":10,"rawMarkdown":"Body","authorUserDisplayName":"Flat User"},{"id":11,"content":"Reply","user":{"displayName":"Nested User","userName":"nested-user"}}]}`)
	var resp MessagesResponse
	if err := json.Unmarshal(payload, &resp); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if len(resp.Comments) != 2 {
		t.Fatalf("unexpected comments: %+v", resp.Comments)
	}
	if resp.Comments[0].AuthorUserDisplayName != "Flat User" {
		t.Fatalf("flat author missing: %+v", resp.Comments[0])
	}
	if resp.Comments[1].User.DisplayName != "Nested User" || resp.Comments[1].User.UserName != "nested-user" {
		t.Fatalf("nested author missing: %+v", resp.Comments[1].User)
	}
}
