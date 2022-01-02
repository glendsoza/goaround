package widgets

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

type LoadingWD struct {
	*tview.TextView
}

// Create and return a new loading widget
func NewLoadingWidget() *LoadingWD {
	return &LoadingWD{tview.NewTextView()}
}

// Display the loading widget
func (lwd *LoadingWD) Load(app *tview.Application, updateFunc func(), doneChan chan int) {
	idx := 0
	dots := []string{".", "..", "...", "...."}
	for {
		select {
		// when work is done we want to stop displaying the loading widget and display the next widget
		case <-doneChan:
			// update the queue to draw the next widget and return
			app.QueueUpdateDraw(updateFunc)
			return
		default:
			if idx >= len(dots) {
				idx = 0
			}
			lwd.SetText(fmt.Sprintf("[#90ee90]Loading %s", dots[idx]))
			idx++
			time.Sleep(time.Second)
			// Draw the loading widget
			app.QueueUpdateDraw(func() {
				app.SetRoot(lwd, true)
			})
		}
	}
}
