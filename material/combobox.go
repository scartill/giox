package material

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/unit"
	giomat "gioui.org/widget/material"
	"github.com/scartill/giox"
)

// ComboStyle holds combobox rendering parameters
type ComboStyle struct {
	theme *giomat.Theme
}

// Combo constructs c ComboStyle
func Combo(theme *giomat.Theme) ComboStyle {
	return ComboStyle{
		theme: theme,
	}
}

// Layout a combobox
func (c ComboStyle) Layout(gtx *layout.Context, widget *giox.Combo) {

	subwidgets := make([]layout.FlexChild, 0)
	if !widget.IsExpanded() {
		for widget.SelectButton().Clicked(gtx) {
			widget.Toggle()
		}

		text := fmt.Sprintf("<%s>", widget.Hint())

		if widget.HasSelected() {
			text = widget.SelectedText()
		}

		subwidgets = append(subwidgets, RigidButton(gtx, c.theme, text, widget.SelectButton()))
	} else {
		N := widget.Len()
		for i := 0; i < N; i++ {
			for widget.Button(i).Clicked(gtx) {
				if err := widget.SelectIndex(i); err != nil {
					fmt.Println("giox error: bad index")
				}
				widget.Toggle()
			}
		}
		
		for i := 0; i < N; i++ {
			subwidgets = append(subwidgets, RigidButton(gtx, c.theme, widget.Item(i), widget.Button(i)))
		}
	}
	
	var inset float32 = 0.0
	if widget.IsExpanded() {
		inset = 10
	}

	layout.Inset{Left: unit.Dp(inset)}.Layout(gtx, func() {
		layout.Flex{Axis: layout.Vertical}.Layout(gtx, subwidgets...)
	})
}
