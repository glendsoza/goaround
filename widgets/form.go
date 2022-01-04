package widgets

import (
	"github.com/rivo/tview"
)

type FormWD struct {
	*tview.Form
}

func NewFormWidget() *FormWD {
	return &FormWD{tview.NewForm()}
}

func (fwd *FormWD) Render() tview.Primitive {
	return fwd
}
