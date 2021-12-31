package main

import (
	"goaround/widgets"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

func main() {
	qwd := widgets.NewQuestionWidget()
	qwd.SetSelectedBackgroundColor(tcell.ColorDarkCyan)
	qwd.ShowSecondaryText(false)
	qwd.SetBorder(true)
	qwd.SetBorderColor(tcell.ColorSnow)
	loading := widgets.NewLoadingWidget()
	loading.SetTitle("[red]Please wait, Quering stack overflow api")
	loading.SetBorder(true)
	loading.SetBorderColor(tcell.ColorSnow)
	loading.SetDynamicColors(true)
	qwd.SetSelectedFunc(func(a int, b, c string, d rune) {
		doneChan := make(chan int)
		awd := widgets.NewAnswerWidget(qwd.GetSelectedQuestion(a))
		awd.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyBackspace2 || event.Key() == tcell.KeyBackspace {
				app.SetRoot(qwd.Render(), true)
				return nil
			}
			return event
		})
		awd.SetBorder(true)
		awd.SetWrap(true)
		awd.SetDynamicColors(true)
		awd.SetBorderColor(tcell.ColorSnow)
		awd.SetToggleHighlights(true)
		go awd.Populate(doneChan)
		go loading.Load(app, func() {
			app.SetRoot(awd.Render(), true)
		}, doneChan)
		app.SetRoot(loading, true)
	})
	doneChan := make(chan int)
	go qwd.Populate(doneChan)
	go loading.Load(app, func() {
		app.SetRoot(qwd.Render(), true)
	}, doneChan)
	if err := app.SetRoot(loading, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
