package material

import (
	"image/color"

	"gioui.org/layout"
	l "gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	giomat "gioui.org/widget/material"
)

// RigidEditor returns layout function for labeled edit field
func RigidEditor(th *giomat.Theme, caption string, hint string, editor *widget.Editor) l.FlexChild {
	editorStyle := giomat.Editor(th, editor, hint)
	editorStyle.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}
	return l.Rigid(func(gtx l.Context) l.Dimensions {
		inset := l.UniformInset(unit.Dp(3))
		return l.Flex{Axis: l.Horizontal}.Layout(gtx,
			l.Rigid(func(gtx l.Context) l.Dimensions {
				return inset.Layout(gtx, func(l.Context) l.Dimensions {
					return giomat.Label(th, unit.Px(16), caption).Layout(gtx)
				})
			}),
			l.Rigid(func(gtx l.Context) l.Dimensions {
				return inset.Layout(gtx, func(gtx l.Context) l.Dimensions {
					return editorStyle.Layout(gtx)
				})
			}))
	})
}

// RigidButton returns layout function for a button with inset
func RigidButton(th *giomat.Theme, caption string, button *widget.Clickable) l.FlexChild {
	inset := l.UniformInset(unit.Dp(3))
	return l.Rigid(func(gtx l.Context) l.Dimensions {
		return inset.Layout(gtx, func(gtx l.Context) l.Dimensions {
			return giomat.Button(th, button, caption).Layout(gtx)
		})
	})
}

// RigidSection returns layout function for a form heading
func RigidSection(th *giomat.Theme, caption string) l.FlexChild {
	inset := l.UniformInset(unit.Dp(3))
	return l.Rigid(func(gtx l.Context) l.Dimensions {
		return inset.Layout(gtx, func(gtx l.Context) l.Dimensions {
			return giomat.H5(th, caption).Layout(gtx)
		})
	})
}

// RigidLabel returns layout function for a regular label
func RigidLabel(th *giomat.Theme, caption string) l.FlexChild {
	return layout.Rigid(func(gtx l.Context) l.Dimensions {
		return giomat.Label(th, unit.Px(16), caption).Layout(gtx)
	})
}

// RigidCheckBox returns layout function for a regular checkbox
func RigidCheckBox(th *giomat.Theme, label string, checkbox *widget.Bool) l.FlexChild {
	return layout.Rigid(func(gtx l.Context) l.Dimensions {
		return giomat.CheckBox(th, checkbox, label).Layout(gtx)
	})
}
