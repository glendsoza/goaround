package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// Add the common parameters to each stack overflow request
func prepareRequest(url, filter string) (*http.Request, url.Values) {
	req, _ := http.NewRequest("GET", url, nil)
	urlQuery := req.URL.Query()
	urlQuery.Add("key", os.Getenv("STACKOVERFLOW_APP_KEY"))
	urlQuery.Add("order", "desc")
	urlQuery.Add("site", "stackoverflow")
	urlQuery.Add("filter", filter)
	pageSize := os.Getenv("STACKOVERFLOW_PAGE_SIZE")
	if pageSize == "" {
		pageSize = "25"
	}
	urlQuery.Add("pagesize", pageSize)

	return req, urlQuery
}

func GetAnswer(qid int) ([]*Answer, error) {
	req, urlQuery := prepareRequest(fmt.Sprintf(STACK_OVERFLOW_ANSWER_URL, qid), "!)xh)am6dFD--YIhimaEuiQq")
	req.URL.RawQuery = urlQuery.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	// in case of error return the error to the caller
	if err != nil {
		return nil, err
	}
	var answerResult AnswerResult
	json.NewDecoder(resp.Body).Decode(&answerResult)
	// update the quota object
	CurrentQuota.QuotaMax = answerResult.QuotaMax
	CurrentQuota.QuotaRemaining = answerResult.QuotaRemaining
	return answerResult.Items, nil
}

func Search() (*SearchResult, error) {
	req, urlQuery := prepareRequest(STACK_OVERFLOW_SEARCH_URL, "!6VClR6PL.AoK9*EK(Zdsdl0uY")
	urlQuery.Add("q", Query)
	urlQuery.Add("sort", "relevance")
	// get the questions with atleast 1 answer
	urlQuery.Add("answers", "1")
	if Tags != "nil" {
		urlQuery.Add("tagged", Tags)
	}
	req.URL.RawQuery = urlQuery.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	// in case of error return the error to the caller
	if err != nil {
		return nil, err
	}
	var searchResult SearchResult
	json.NewDecoder(resp.Body).Decode(&searchResult)
	// update the quota object
	CurrentQuota.QuotaMax = searchResult.QuotaMax
	CurrentQuota.QuotaRemaining = searchResult.QuotaRemaining
	return &searchResult, nil

}
