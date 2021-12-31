package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func prepareRequest(url, filter string) (*http.Request, url.Values) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Something went wrong while querying stackexchange api")
	}
	urlQuery := req.URL.Query()
	urlQuery.Add("key", os.Getenv("STACKOVERFLOW_APP_KEY"))
	urlQuery.Add("order", "desc")
	urlQuery.Add("sort", "activity")
	urlQuery.Add("site", "stackoverflow")
	urlQuery.Add("filter", filter)
	pageSize := os.Getenv("STACKOVERFLOW_PAGE_SIZE")
	if pageSize == "" {
		pageSize = "25"
	}
	urlQuery.Add("pagesize", pageSize)

	return req, urlQuery
}

func GetAnswer(qid int) []*Answer {
	req, urlQuery := prepareRequest(fmt.Sprintf(STACK_OVERFLOW_ANSWER_URL, qid), "!)xh)am6dFD--YIhimaEuiQq")
	req.URL.RawQuery = urlQuery.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		log.Fatal("Something went wrong while querying stackexchange api")
	}
	var answerResult AnswerResult
	json.NewDecoder(resp.Body).Decode(&answerResult)
	CurrentQuota.QuotaMax = answerResult.QuotaMax
	CurrentQuota.QuotaRemaining = answerResult.QuotaRemaining
	return answerResult.Items
}

func Search(q string) *SearchResult {
	req, urlQuery := prepareRequest(STACK_OVERFLOW_SEARCH_URL, "!6VClR6PL.AoK9*EK(Zdsdl0uY")
	urlQuery.Add("q", q)
	req.URL.RawQuery = urlQuery.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		log.Fatal("Something went wrong while querying stackexchange api")
	}
	var searchResult SearchResult
	json.NewDecoder(resp.Body).Decode(&searchResult)
	CurrentQuota.QuotaMax = searchResult.QuotaMax
	CurrentQuota.QuotaRemaining = searchResult.QuotaRemaining
	return &searchResult

}
