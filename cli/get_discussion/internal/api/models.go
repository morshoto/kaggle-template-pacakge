package api

type TopicResponse struct {
	ForumTopic struct {
		Name                  string `json:"name"`
		URL                   string `json:"url"`
		AuthorUserDisplayName string `json:"authorUserDisplayName"`
		AuthorUserName        string `json:"authorUserName"`
		TotalMessages         *int   `json:"totalMessages"`
		PostDate              string `json:"postDate"`
		FirstMessageID        int    `json:"firstMessageId"`
	} `json:"forumTopic"`
}

type MessagesResponse struct {
	Comments []struct {
		ID          int    `json:"id"`
		RawMarkdown string `json:"rawMarkdown"`
		Content     string `json:"content"`
	} `json:"comments"`
}

type CompetitionResponse struct {
	ForumID *int `json:"forumId"`
}

type TopicListResponse struct {
	Count  int `json:"count"`
	Topics []struct {
		TopicURL string `json:"topicUrl"`
		URL      string `json:"url"`
	} `json:"topics"`
}
