package widgets

import (
	"bytes"
	"goaround/api"
	"goaround/constants"
	gwt "goaround/templates"
	"goaround/utils"
	"sort"
	"text/template"

	"github.com/rivo/tview"
)

var answerCache = make(map[int]*api.AnswerResult)

type AnswerWD struct {
	*tview.TextView
	question       *api.Question
	answerTemplate *template.Template
}

// Create and return a new answer widget
func NewAnswerWidget() *AnswerWD {
	awd := &AnswerWD{tview.NewTextView(), nil, nil}
	awd.initializeTemplate()
	return awd
}

func (awd *AnswerWD) IsTemplateInitialized() bool {
	return awd.answerTemplate != nil
}

func (awd *AnswerWD) initializeTemplate() {
	// Initialize the template for the answer
	t, err := template.New("AnswerTemplate").Funcs(template.FuncMap{
		"BeautifyHtmlText":  utils.BeautifyHtmlText,
		"GetDateDiffInDays": utils.GetDateDiffInDays,
		"Add": func(i, j int) int {
			return i + j
		},
	}).Parse(gwt.AnswerTemplate)
	// If template initialization fails return the error to the caller
	if err == nil {
		awd.answerTemplate = t
	}
}

func (awd *AnswerWD) SetQuestion(question *api.Question) {
	awd.question = question
}

// Get the answers for the given question
func (awd *AnswerWD) getAnswer() (*api.AnswerResult, error) {
	// Check if answer is already cached
	data, ok := answerCache[awd.question.QuestionID]
	if !ok {
		// If not cached, fetch the answers from the API
		data, err := api.GetAnswer(awd.question.QuestionID)
		// In case of error call the error handler
		if err != nil {
			return nil, err
		}
		// sort the answers by accepted field
		sort.Slice(data.Items, func(i int, j int) bool {
			return data.Items[i].IsAccepted
		})
		answerCache[awd.question.QuestionID] = data
		return data, nil
	}
	return data, nil
}

// Populates the answer widget
func (awd *AnswerWD) Populate(doneChan chan int) {
	awd.Clear()
	buf := &bytes.Buffer{}
	// get the answers
	answers, err := awd.getAnswer()
	if err != nil {
		awd.SetText("[red]Something went wrong while calling api[-]")
		doneChan <- 1
		return
	}
	err = awd.answerTemplate.Execute(buf, struct {
		Question        *api.Question
		SeperatorString string
		Answers         []*api.Answer
		QuestionURL     string
	}{Question: awd.question,
		SeperatorString: utils.GenerateSeperatorString(25),
		Answers:         answers.Items,
		QuestionURL:     constants.STACK_OVERFLOW_UI_QUESTION_URL})
	// In case of error call the error handler
	if err != nil {
		buf.WriteString("[red]Something went wrong while rendering the anwer[-]")
	}
	// Replace 2 or more new lines with a single new line
	awd.SetText(constants.REPLACE_MULTIPLE_NEW_LINE_REGEX.ReplaceAllString(buf.String(), "\n\n"))
	// send a signal to indicate answers have been loaded
	doneChan <- 1
}

func (awd *AnswerWD) Render() tview.Primitive {
	return awd
}
