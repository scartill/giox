package material

import (
	"fmt"
	"image"

	"gioui.org/layout"
	"gioui.org/io/pointer"
	"gioui.org/unit"
	"gioui.org/op"
	giomat "gioui.org/widget/material"
	"github.com/scartill/giox"
)

// ComboStyle holds combobox rendering parameters
type ComboStyle struct {
	size unit.Value
	theme *giomat.Theme
}

// Combo constructs c ComboStyle
func Combo(theme *giomat.Theme, size unit.Value) ComboStyle {
	return ComboStyle{
		theme: theme,
		size: size,
	}
}

// Layout a combobox
func (c ComboStyle) Layout(gtx *layout.Context, widget *giox.Combo) {

	for _, e := range gtx.Events(widget) {
		if e, ok := e.(pointer.Event); ok {
			if e.Type == pointer.Press {
				widget.Toggle()
			}
		}
	}

	if !widget.IsExpanded() {
		text := fmt.Sprintf("<%s>", widget.Hint())

		if widget.HasSelected() {
			text = widget.SelectedText()
		}

		giomat.Label(c.theme, c.size, text).Layout(gtx)

	} else {
		items := widget.Items()
		labels := make([]layout.FlexChild, len(items))
		for i, val := range items {
			labels[i] = RigidLabel(gtx, c.theme, val)
		}
		layout.Flex{Axis: layout.Vertical}.Layout(gtx, labels...)
	}

	var st op.StackOp
	st.Push(gtx.Ops)
	clickRect := image.Rectangle{Max: gtx.Dimensions.Size}
	pointer.Rect(clickRect).Add(gtx.Ops)
	pointer.InputOp{Key: widget}.Add(gtx.Ops)
	st.Pop()
}
