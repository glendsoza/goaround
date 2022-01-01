package widgets

import (
	"bytes"
	"goaround/api"
	gwt "goaround/templates"
	"goaround/utils"
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

func NewAnswerWidget(question *api.Question) (*AnswerWD, error) {
	t, err := template.New("AnswerTemplate").Funcs(template.FuncMap{
		"BeautifyHtmlText":  utils.BeautifyHtmlText,
		"GetDateDiffInDays": utils.GetDateDiffInDays,
		"Add": func(i, j int) int {
			return i + j
		},
	}).Parse(gwt.AnswerTemplate)
	if err != nil {
		return nil, err
	}
	return &AnswerWD{tview.NewTextView(), question, t}, nil
}

func (awd *AnswerWD) GetAnswer(errorHandler func(error)) []*api.Answer {
	data, ok := answerCache[awd.question.QuestionID]
	if !ok {
		data, err := api.GetAnswer(awd.question.QuestionID)
		if err != nil {
			errorHandler(err)
		}
		sort.Slice(data, func(i int, j int) bool {
			return data[i].IsAccepted
		})
		answerCache[awd.question.QuestionID] = data
		return data
	}
	return data
}

func (awd *AnswerWD) Populate(doneChan chan int, errorHandler func(error)) {
	buf := &bytes.Buffer{}
	err := awd.answerTemplate.Execute(buf, struct {
		Question        *api.Question
		SeperatorString string
		Answers         []*api.Answer
	}{Question: awd.question,
		SeperatorString: utils.GenerateSeperatorString(25),
		Answers:         awd.GetAnswer(errorHandler)})
	if err != nil {
		errorHandler(err)
	}
	awd.SetText(REPLACE_MULTIPLE_NEW_LINE_REGEX.ReplaceAllString(buf.String(), "\n\n"))
	doneChan <- 1
}
