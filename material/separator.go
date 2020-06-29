package material

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	l "gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	giomat "gioui.org/widget/material"
	"github.com/scartill/giox"
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

		d := image.Point{X: int(width), Y: gtx.Px(height)}
		dr := f32.Rectangle{
			Max: f32.Point{X: float32(d.X), Y: float32(d.Y)},
		}

		clip.Rect{
			Rect: f32.Rectangle{Max: f32.Point{X: width, Y: float32(gtx.Px(height))}},
		}.Op(gtx.Ops).Add(gtx.Ops)

		paint.ColorOp{Color: color.RGBA{
			R: 128, G: 128, B: 128, A: 255,
		}}.Add(gtx.Ops)
		paint.PaintOp{Rect: dr}.Add(gtx.Ops)

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
