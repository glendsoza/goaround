package ui

import (
	"goaround/constants"
	"goaround/widgets"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

type Manager struct {
	qwd *widgets.QuestionWD
	awd *widgets.AnswerWD
	lwd *widgets.LoadingWD
	fwd *widgets.FormWD
	wwd *widgets.WrapperWD
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

func initWrapperWidget() *widgets.WrapperWD {
	return widgets.NewWrapperWD()
}

func NewManager() *Manager {
	// Initialize all the widgets
	m := &Manager{
		qwd: initQuestionWD(),
		awd: initAnswerWD(),
		lwd: initLoadingWD(),
		fwd: initFormWidget(),
		wwd: initWrapperWidget(),
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
			m.displayForm(func() {
				m.renderAnswer()
			})
			return nil
		}
		return event
	})
}

func (m *Manager) setQuestionInputCapture() {
	m.qwd.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlR:
			{
				m.displayForm(func() {
					m.renderQuestion()
				})
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

func (m *Manager) waitForQuestionLoad() {
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

func (m *Manager) render(primitive tview.Primitive) {
	app.SetRoot(primitive, true)
}

func (m *Manager) renderQuestion() {
	m.wwd.AddItem(m.qwd.Render())
	m.wwd.SetText(constants.QUESTION_FOOTER)
	m.render(m.wwd.Render())
}

func (m *Manager) renderAnswer() {
	m.wwd.AddItem(m.awd.Render())
	m.wwd.SetText(constants.ANSWER_FOOTER)
	m.render(m.wwd.Render())
}

func (m *Manager) renderLoading() {
	app.SetRoot(m.lwd.Render(), true)
}

func (m *Manager) renderForm() {
	m.wwd.AddItem(m.fwd.Render())
	m.wwd.SetText(constants.FORM_FOOTER)
	m.render(m.wwd.Render())
}

func (m *Manager) displayForm(onReturnFunc func()) {
	m.fwd.Clear(true)
	m.fwd.AddInputField("Query", "", 1000, nil, nil).
		AddInputField("Tags", "", 1000, nil, nil).
		AddButton("Submit", func() {
			m.qwd.SetQuery(m.fwd.GetFormItem(0).(*tview.InputField).GetText())
			m.qwd.SetTags(m.fwd.GetFormItem(1).(*tview.InputField).GetText())
			m.waitForQuestionLoad()
		})
	m.fwd.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlR:
			{
				onReturnFunc()
				return nil
			}
		}
		return event
	})
	m.renderForm()
}

func (m *Manager) Run() error {
	m.setQuestionInputCapture()
	m.setAnswerInputCapture()
	m.setSelectedQuestionHandler()
	m.waitForQuestionLoad()
	if err := app.SetRoot(m.lwd, true).EnableMouse(false).Run(); err != nil {
		return err
	}
	return nil
}
