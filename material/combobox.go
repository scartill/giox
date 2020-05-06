package material

import (
	"gioui.org/layout"
	"gioui.org/unit"
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
	giomat.Label(c.theme, unit.Px(16), widget.Selected()).Layout(gtx)
}
