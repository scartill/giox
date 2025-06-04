package material

import (
	"image"

	"gioui.org/layout"
	l "gioui.org/layout"
	"gioui.org/unit"
	giomat "gioui.org/widget/material"
	"github.com/nkrul/giox"
)

// SeparatorStyle defines material rendering parameters for separator
type SeparatorStyle struct {
	widget *giox.Separator
	theme  *giomat.Theme
}

// Separator creates a material separator widget
func Separator(theme *giomat.Theme, widget *giox.Separator) SeparatorStyle {
	return SeparatorStyle{
		widget: widget,
		theme:  theme,
	}
}

// Layout a separator
func (ss SeparatorStyle) Layout(gtx l.Context) l.Dimensions {

	shader := func(width float32) layout.Dimensions {
		height := unit.Dp(2)

		d := image.Point{X: int(width), Y: gtx.Dp(height)}

		return layout.Dimensions{Size: d}
	}

	realWidth := float32(gtx.Constraints.Max.X)

	gtx.Constraints.Min.X = int(realWidth)

	return layout.Stack{Alignment: layout.W}.Layout(gtx,
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return shader(realWidth)
		}),
	)
}
