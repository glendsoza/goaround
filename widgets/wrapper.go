package widgets

import "github.com/rivo/tview"

type WrapperWD struct {
	*tview.Flex
	*tview.TextView
}

func NewWrapperWD() *WrapperWD {
	return &WrapperWD{tview.NewFlex().SetDirection(tview.FlexRow),
		tview.NewTextView().SetDynamicColors(true)}
}

// Add the given primitive as a main item
func (wwd *WrapperWD) AddItem(item tview.Primitive) {
	// Clear any previously set item
	wwd.Flex.Clear()
	// Add the text view as well
	wwd.Flex.AddItem(item, 0, 1, true).
		AddItem(wwd.TextView, 1, 1, false)
}

func (wwd *WrapperWD) SetText(text string) {
	// At the time of rendering text set via SetText method will be displyed on the terminal
	// Clears any previously set text
	wwd.TextView.SetText(text)
}

func (wwd *WrapperWD) Render() tview.Primitive {
	return wwd.Flex
}
