package api

type ErrorResponse struct {
	ErrorID      int    `json:"error_id"`
	Backoff      int    `json:"backoff"`
	ErrorMessage string `json:"error_message"`
	ErrorName    string `json:"error_name"`
}
type APIQuota struct {
	HasMore        bool `json:"has_more"`
	QuotaMax       int  `json:"quota_max"`
	QuotaRemaining int  `json:"quota_remaining"`
}
type Answer struct {
	DownVoteCount    int    `json:"down_vote_count"`
	UpVoteCount      int    `json:"up_vote_count"`
	IsAccepted       bool   `json:"is_accepted"`
	LastActivityDate int    `json:"last_activity_date"`
	ShareLink        string `json:"share_link"`
	Body             string `json:"body"`
	AnswerID         string `json:"answer_id"`
}

type Question struct {
	UpVoteCount      int    `json:"up_vote_count"`
	AnswerCount      int    `json:"answer_count"`
	LastActivityDate int    `json:"last_activity_date"`
	CreationDate     int    `json:"creation_date"`
	QuestionID       int    `json:"question_id"`
	Body             string `json:"body"`
	BodyMarkdown     string `json:"body_markdown"`
	Title            string `json:"title"`
}

type Base struct {
	APIQuota
	ErrorResponse
}

type SearchResult struct {
	Items []*Question
	Base
}

type AnswerResult struct {
	Items []*Answer
	Base
}
