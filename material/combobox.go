package material

import (
	"fmt"

	"gioui.org/layout"
	l "gioui.org/layout"
	"gioui.org/unit"
	giomat "gioui.org/widget/material"
	"github.com/nkrul/giox"
)

// ComboStyle holds combobox rendering parameters
type ComboStyle struct {
	widget *giox.Combo
	theme  *giomat.Theme
}

// Combo constructs c ComboStyle
func Combo(theme *giomat.Theme, widget *giox.Combo) ComboStyle {
	return ComboStyle{
		widget: widget,
		theme:  theme,
	}
}

// Layout a combobox
func (c ComboStyle) Layout(gtx l.Context) l.Dimensions {

	subwidgets := make([]l.FlexChild, 0)
	if !c.widget.IsExpanded() {
		for c.widget.SelectButton().Clicked(gtx) {
			c.widget.Toggle()
		}

		text := fmt.Sprintf("<%s>", c.widget.Hint())

		if c.widget.HasSelected() {
			text = c.widget.SelectedText()
		}

		subwidgets = append(subwidgets, RigidButton(c.theme, text, c.widget.SelectButton()))
	} else {
		N := c.widget.Len()
		for i := 0; i < N; i++ {
			for c.widget.Button(i).Clicked(gtx) {
				if err := c.widget.SelectIndex(i); err != nil {
					fmt.Println("giox error: bad index")
				}
				c.widget.Toggle()
			}
		}

		for i := 0; i < N; i++ {
			subwidgets = append(subwidgets, RigidButton(c.theme, c.widget.Item(i), c.widget.Button(i)))
		}
	}

	var inset float32 = 0.0
	if c.widget.IsExpanded() {
		inset = 10
	}

	return l.Inset{Left: unit.Dp(inset)}.Layout(gtx, func(layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx, subwidgets...)
	})
}
