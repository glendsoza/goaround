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

func (wwd *WrapperWD) AddItem(item tview.Primitive) {
	wwd.Flex.Clear()
	wwd.Flex.AddItem(item, 0, 1, true).
		AddItem(wwd.TextView, 1, 1, false)
}

func (wwd *WrapperWD) SetText(text string) {
	wwd.TextView.SetText(text)
}

func (wwd *WrapperWD) Render() tview.Primitive {
	return wwd.Flex
}
