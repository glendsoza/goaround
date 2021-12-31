package widgets

import (
	"bytes"
	"fmt"
	"goaround/api"
	gwt "goaround/templates"
	"goaround/utils"
	"log"
	"os"
	"regexp"
	"sort"
	"text/template"

	"github.com/rivo/tview"
)

var REPLACE_MULTIPLE_NEW_LINE_REGEX = regexp.MustCompile("(\n\n)+")
var answerCache = make(map[int][]*api.Answer)

type AnswerWD struct {
	*tview.TextView
	question       *api.Question
	answerTemplate *template.Template
}

func NewAnswerWidget(question *api.Question) *AnswerWD {
	t, _ := template.New("AnswerTemplate").Funcs(template.FuncMap{
		"BeautifyHtmlText":  utils.BeautifyHtmlText,
		"GetDateDiffInDays": utils.GetDateDiffInDays,
	}).Parse(gwt.AnswerTemplate)
	return &AnswerWD{tview.NewTextView(), question, t}
}

func (awd *AnswerWD) Populate(doneChan chan int) {
	buf := &bytes.Buffer{}
	answers, ok := answerCache[awd.question.QuestionID]
	if !ok {
		answers = api.GetAnswer(awd.question.QuestionID)
		sort.Slice(answers, func(i int, j int) bool {
			return answers[i].IsAccepted
		})
		answerCache[awd.question.QuestionID] = answers
	}
	err := awd.answerTemplate.Execute(buf, struct {
		Question        *api.Question
		SeperatorString string
		Answers         []*api.Answer
	}{Question: awd.question, SeperatorString: utils.GenerateSeperatorString(25), Answers: answers})
	if err != nil {
		log.Println(err)
		log.Fatal("Something went wrong while rendering answers")
	}
	awd.SetText(REPLACE_MULTIPLE_NEW_LINE_REGEX.ReplaceAllString(buf.String(), "\n\n"))
	doneChan <- 1
}

func (awd *AnswerWD) Render() *AnswerWD {
	usingKey := "No"
	if os.Getenv("STACKOVERFLOW_APP_KEY") != "" {
		usingKey = "Yes"
	}
	awd.SetTitle(fmt.Sprintf("[red]Quota Max : %d | Quota Remaining : %d | Using Key : %s [-]",
		api.CurrentQuota.QuotaMax,
		api.CurrentQuota.QuotaRemaining, usingKey))
	return awd
}
