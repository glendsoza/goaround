package ui

import (
	"goaround/widgets"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

// Initialize the loading widget
func initLoading() *widgets.LoadingWD {
	lwd := widgets.NewLoadingWidget()
	lwd.SetTitle("[red]Please wait, Querying stack overflow api")
	lwd.SetBorder(true)
	lwd.SetBorderColor(tcell.ColorSnow)
	lwd.SetDynamicColors(true)
	return lwd
}

// Initialize the question widget
func initQuestion() *widgets.QuestionWD {
	qwd := widgets.NewQuestionWidget()
	qwd.SetSelectedBackgroundColor(tcell.ColorDarkCyan)
	qwd.ShowSecondaryText(false)
	qwd.SetBorder(true)
	qwd.SetBorderColor(tcell.ColorSnow)
	return qwd
}

func InIt() {
	qwd := initQuestion()
	lwd := initLoading()
	errorHandler := func(err error) {
		app.Stop()
		log.Fatal(err)
	}
	qwd.SetSelectedFunc(func(a int, b, c string, d rune) {
		// When the questions are not loaded secondary text of the question will be set to error
		// in this case we simply want to return
		if c == "error" {
			return
		}
		doneChan := make(chan int)
		awd, err := widgets.NewAnswerWidget(qwd.GetSelectedQuestion(a))
		if err != nil {
			errorHandler(err)
		}
		awd.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyBackspace2 || event.Key() == tcell.KeyBackspace {
				app.SetRoot(qwd.Render(), true)
				return nil
			}
			return event
		})
		awd.SetWrap(true)
		awd.SetDynamicColors(true)
		awd.SetBorderColor(tcell.ColorSnow)
		awd.SetToggleHighlights(true)
		// call the go routine to populate the answers
		go awd.Populate(doneChan, errorHandler)
		go lwd.Load(app, func() {
			app.SetRoot(awd, true)
		}, doneChan)
		app.SetRoot(lwd, true)
	})
	doneChan := make(chan int)
	// go the go routine to populate questions
	go qwd.Populate(doneChan, errorHandler)
	go lwd.Load(app, func() {
		app.SetRoot(qwd.Render(), true)
	}, doneChan)
	if err := app.SetRoot(lwd, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}
}
