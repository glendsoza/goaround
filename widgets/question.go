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
}

func NewQuestionWidget() *QuestionWD {
	return &QuestionWD{tview.NewList(), make(map[int]*api.Question)}
}

func (qwd *QuestionWD) Populate(doneChan chan int) {
	result := api.Search("all go routines are asleep")
	if result.ErrorID == 400 {
		qwd.AddItem("Invalid key/page size supplied", "", '0', nil)
	} else {
		for idx, data := range result.Items {
			data.Title = html.UnescapeString(data.Title)
			qwd.questionMapping[idx] = data
			qwd.AddItem(fmt.Sprintf("[yellow](%d)[-] %s (%d Answers)",
				idx,
				data.Title,
				data.AnswerCount), "", 0, nil)
		}
	}

	doneChan <- 1
}
func (qwd *QuestionWD) Render() *QuestionWD {
	if qwd.GetItemCount() == 0 {
		qwd.AddItem("No Data To Display", "", '0', nil)
	}
	usingKey := "No"
	if os.Getenv("STACKOVERFLOW_APP_KEY") != "" {
		usingKey = "Yes"
	}
	qwd.SetTitle(fmt.Sprintf("[red]Quota Max : %d | Quota Remaining : %d | Using Key : %s [-]",
		api.CurrentQuota.QuotaMax,
		api.CurrentQuota.QuotaRemaining, usingKey))
	return qwd
}

func (qwd *QuestionWD) GetSelectedQuestion(idx int) *api.Question {
	return qwd.questionMapping[idx]
}
