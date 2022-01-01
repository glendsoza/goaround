package main

import (
	"flag"
	"goaround/widgets"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()
var query string

func initLoading() *widgets.Loading {
	loading := widgets.NewLoadingWidget()
	loading.SetTitle("[red]Please wait, Querying stack overflow api")
	loading.SetBorder(true)
	loading.SetBorderColor(tcell.ColorSnow)
	loading.SetDynamicColors(true)
	return loading
}

func initQuestion() *widgets.QuestionWD {
	qwd := widgets.NewQuestionWidget()
	qwd.SetSelectedBackgroundColor(tcell.ColorDarkCyan)
	qwd.ShowSecondaryText(false)
	qwd.SetBorder(true)
	qwd.SetBorderColor(tcell.ColorSnow)
	return qwd
}

func main() {
	flag.StringVar(&query, "q", "", "Query to search")
	flag.Parse()
	if query == "" {
		log.Fatal("Please pass the query with -q option")
	}
	qwd := initQuestion()
	loading := initLoading()
	errorHandler := func(err error) {
		app.Stop()
		log.Fatal(err)
	}
	qwd.SetSelectedFunc(func(a int, b, c string, d rune) {
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
		go awd.Populate(doneChan, errorHandler)
		go loading.Load(app, func() {
			app.SetRoot(awd, true)
		}, doneChan)
		app.SetRoot(loading, true)
	})
	doneChan := make(chan int)
	go qwd.Populate(doneChan, query, errorHandler)
	go loading.Load(app, func() {
		app.SetRoot(qwd.Render(), true)
	}, doneChan)
	if err := app.SetRoot(loading, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}
}
