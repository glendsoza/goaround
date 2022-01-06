package widgets

import (
	"fmt"
	"goaround/api"
	"html"
	"os"

	"github.com/rivo/tview"
)

type QuestionWD struct {
	*tview.List
	questionMapping map[int]*api.Question
	query           string
	tags            string
}

func (qwd *QuestionWD) SetQuery(query string) {
	qwd.query = query
}

func (qwd *QuestionWD) SetTags(tags string) {
	qwd.tags = tags
}

func (qwd *QuestionWD) GetQuery() string {
	return qwd.query
}

// Create and return a new question widget
func NewQuestionWidget() *QuestionWD {
	return &QuestionWD{tview.NewList(), make(map[int]*api.Question), "", ""}
}

// Populates the question widget
func (qwd *QuestionWD) Populate(doneChan chan int) {
	qwd.Clear()
	result, err := api.Search(qwd.query, qwd.tags)
	// In case of error call the error handler
	if err != nil {
		qwd.AddItem("Something went wrong while calling api", "error", '0', nil)
		doneChan <- 1
		return
	}
	// if api returns error set the secondary text to error
	if result.ErrorID == 400 {
		qwd.AddItem("Invalid key/page size supplied", "error", '0', nil)
	} else {
		for idx, data := range result.Items {
			data.Title = html.UnescapeString(data.Title)
			// create the mapping between question id and question
			qwd.questionMapping[idx] = data
			qwd.AddItem(fmt.Sprintf("[yellow](%d)[-] %s (%d Answers)",
				idx,
				data.Title,
				data.AnswerCount), "", 0, nil)
		}
	}
	// send a signal to indicate quesiton has been loaded
	doneChan <- 1
}

// Wrapper before rendering the widget
func (qwd *QuestionWD) Render() tview.Primitive {
	// if no question are found then set the secondary text to error to suppress on select function
	if qwd.GetItemCount() == 0 {
		qwd.AddItem("No Data To Display", "error", '0', nil)
	}
	usingKey := "No"
	if os.Getenv("STACKOVERFLOW_APP_KEY") != "" {
		usingKey = "Yes"
	}
	// set the quotas and key in the title
	remainingQuota, maxQuota := api.GetCurrentQuota()
	qwd.SetTitle(fmt.Sprintf("[red]Quota Max : %d | Quota Remaining : %d | Using Key : %s [-]",
		maxQuota,
		remainingQuota, usingKey))
	return qwd
}

// Returns the question object based on the selected question
func (qwd *QuestionWD) GetSelectedQuestion(idx int) *api.Question {
	return qwd.questionMapping[idx]
}
