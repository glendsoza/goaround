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

// Create and return a new answer widget
func NewAnswerWidget(question *api.Question) (*AnswerWD, error) {
	// Initialize the template for the answer
	t, err := template.New("AnswerTemplate").Funcs(template.FuncMap{
		"BeautifyHtmlText":  utils.BeautifyHtmlText,
		"GetDateDiffInDays": utils.GetDateDiffInDays,
		"Add": func(i, j int) int {
			return i + j
		},
	}).Parse(gwt.AnswerTemplate)
	// If template initialization fails return the error to the caller
	if err != nil {
		return nil, err
	}
	return &AnswerWD{tview.NewTextView(), question, t}, nil
}

// Get the answers for the given question
func (awd *AnswerWD) GetAnswer(errorHandler func(error)) []*api.Answer {
	// Check if answer is already cached
	data, ok := answerCache[awd.question.QuestionID]
	if !ok {
		// If not cached, fetch the answers from the API
		data, err := api.GetAnswer(awd.question.QuestionID)
		// In case of error call the error handler
		if err != nil {
			errorHandler(err)
		}
		// sort the answers by accepted field
		sort.Slice(data, func(i int, j int) bool {
			return data[i].IsAccepted
		})
		answerCache[awd.question.QuestionID] = data
		return data
	}
	return data
}

// Populates the answer widget
func (awd *AnswerWD) Populate(doneChan chan int, errorHandler func(error)) {
	buf := &bytes.Buffer{}
	err := awd.answerTemplate.Execute(buf, struct {
		Question        *api.Question
		SeperatorString string
		Answers         []*api.Answer
	}{Question: awd.question,
		SeperatorString: utils.GenerateSeperatorString(25),
		Answers:         awd.GetAnswer(errorHandler)})
	// In case of error call the error handler
	if err != nil {
		errorHandler(err)
	}
	// Replace 2 or more new lines with a single new line
	awd.SetText(REPLACE_MULTIPLE_NEW_LINE_REGEX.ReplaceAllString(buf.String(), "\n\n"))
	// send a signal to indicate answers have been loaded
	doneChan <- 1
}
