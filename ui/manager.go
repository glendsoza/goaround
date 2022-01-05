package ui

import (
	"goaround/widgets"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

type Renderable interface {
	Render() tview.Primitive
}

type Manager struct {
	qwd *widgets.QuestionWD
	awd *widgets.AnswerWD
	lwd *widgets.LoadingWD
	fwd *widgets.FormWD
}

func initQuestionWD() *widgets.QuestionWD {
	qwd := widgets.NewQuestionWidget()
	qwd.SetSelectedBackgroundColor(tcell.ColorDarkCyan)
	qwd.ShowSecondaryText(false)
	qwd.SetBorder(true)
	qwd.SetBorderColor(tcell.ColorSnow)
	return qwd
}

func initLoadingWD() *widgets.LoadingWD {
	lwd := widgets.NewLoadingWidget()
	lwd.SetTitle("[red]Please wait, Querying stack overflow api")
	lwd.SetBorder(true)
	lwd.SetBorderColor(tcell.ColorSnow)
	lwd.SetDynamicColors(true)
	return lwd
}

func initAnswerWD() *widgets.AnswerWD {
	awd := widgets.NewAnswerWidget()
	awd.SetWrap(true)
	awd.SetDynamicColors(true)
	awd.SetBorderColor(tcell.ColorSnow)
	awd.SetToggleHighlights(true)
	return awd

}

func initFormWidget() *widgets.FormWD {
	return widgets.NewFormWidget()
}

func NewManager() *Manager {
	// Initialize all the widgets
	m := &Manager{
		qwd: initQuestionWD(),
		awd: initAnswerWD(),
		lwd: initLoadingWD(),
		fwd: initFormWidget(),
	}
	return m
}

func (m *Manager) SetQuestionQuery(query string) {
	m.qwd.SetQuery(query)
}

func (m *Manager) SetQuestionTags(tags string) {
	m.qwd.SetTags(tags)
}

func (m *Manager) setAnswerInputCapture() {
	m.awd.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyBackspace2, tcell.KeyBackspace:
			if event.Key() == tcell.KeyBackspace2 || event.Key() == tcell.KeyBackspace {
				m.renderQuestion()
				return nil
			}
		case tcell.KeyCtrlR:
			m.displayForm(m.awd)
			return nil
		}
		return event
	})
}

// Change this to populate
func (m *Manager) displayForm(onCancelPrimitive Renderable) {
	m.fwd.Clear(true)
	m.fwd.AddInputField("Query", "", 1000, nil, nil).
		AddInputField("Tags", "", 1000, nil, nil).
		AddButton("Submit", func() {
			m.qwd.SetQuery(m.fwd.GetFormItem(0).(*tview.InputField).GetText())
			m.qwd.SetTags(m.fwd.GetFormItem(1).(*tview.InputField).GetText())
			m.loadQuestions()
		})
	m.fwd.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlR:
			{
				app.SetRoot(onCancelPrimitive.Render(), true)
				return nil
			}
		}
		return event
	})
	m.renderForm()
}

func (m *Manager) setQuestionInputCapture() {
	m.qwd.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlR:
			{
				m.displayForm(m.qwd)
				return nil
			}
		}
		return event
	})
}

func (m *Manager) waitForAnswerLoad() {
	doneChan := make(chan int)
	go m.awd.Populate(doneChan)
	go m.lwd.Load(app, func() {
		m.renderAnswer()
	}, doneChan)
	m.renderLoading()
}

func (m *Manager) loadQuestions() {
	doneChan := make(chan int)
	// go the go routine to populate questions
	go m.qwd.Populate(doneChan)
	go m.lwd.Load(app, func() {
		m.renderQuestion()
	}, doneChan)
}

func (m *Manager) setSelectedQuestionHandler() {
	m.qwd.SetSelectedFunc(func(idx int, b, c string, d rune) {
		// When the questions are not loaded secondary text of the question will be set to error
		// in this case we simply want to return
		if c == "error" {
			return
		}
		if !m.awd.IsTemplateInitialized() {
			m.qwd.Clear()
			m.qwd.AddItem("Something went wrong while initializing the answer", "error", '0', nil)
			return
		}
		m.awd.SetQuestion(m.qwd.GetSelectedQuestion(idx))
		// call the go routine to populate the answers
		m.waitForAnswerLoad()
	})
}

func (m *Manager) renderQuestion() {
	app.SetRoot(m.qwd.Render(), true)
}

func (m *Manager) renderAnswer() {
	app.SetRoot(m.awd.Render(), true)

}

func (m *Manager) renderLoading() {
	app.SetRoot(m.lwd.Render(), true)
}

func (m *Manager) renderForm() {
	app.SetRoot(m.fwd.Render(), true)
}

func (m *Manager) Run() error {
	m.setQuestionInputCapture()
	m.setAnswerInputCapture()
	m.setSelectedQuestionHandler()
	m.loadQuestions()
	if err := app.SetRoot(m.lwd, true).EnableMouse(false).Run(); err != nil {
		return err
	}
	return nil
}
