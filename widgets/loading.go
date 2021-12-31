package widgets

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

type Loading struct {
	*tview.TextView
}

func NewLoadingWidget() *Loading {
	return &Loading{tview.NewTextView()}
}

func (l *Loading) Load(app *tview.Application, updateFunc func(), doneChan chan int) {
	idx := 0
	dots := []string{".", "..", "...", "...."}
	for {
		select {
		case <-doneChan:
			app.QueueUpdateDraw(updateFunc)
			return
		default:
			if idx >= len(dots) {
				idx = 0
			}
			l.SetText(fmt.Sprintf("[#90ee90]Loading %s", dots[idx]))
			idx++
			time.Sleep(time.Second)
			app.QueueUpdateDraw(func() {
				app.SetRoot(l, true)
			})
		}
	}
}
