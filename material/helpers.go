package material

import (
    "image/color"

    "gioui.org/layout"
    "gioui.org/unit"
    "gioui.org/widget"
    giomat "gioui.org/widget/material"
)

// RigidEditor returns layout function for labeled edit field
func RigidEditor(gtx *layout.Context, th *giomat.Theme, caption string, hint string, editor *widget.Editor) layout.FlexChild {
    editorStyle := giomat.Editor(th, hint)
   editorStyle.Color = color.RGBA{R:0, G:0, B:255, A:255}
    return layout.Rigid(func() {
        inset := layout.UniformInset(unit.Dp(3))
        layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
            layout.Rigid(func() {
                inset.Layout(gtx, func() {
                    giomat.Label(th, unit.Px(16), caption).Layout(gtx)
                })
            }),
            layout.Rigid(func() {
                inset.Layout(gtx, func() {
                    editorStyle.Layout(gtx, editor)
                })
            }))
    })
}

// RigidButton returns layout function for a button with inset
func RigidButton(gtx *layout.Context, th *giomat.Theme, caption string, button *widget.Button) layout.FlexChild {
    inset := layout.UniformInset(unit.Dp(3))
    return layout.Rigid(func() {
        inset.Layout(gtx, func() {
            giomat.Button(th, caption).Layout(gtx, button)
        })
    })
}

// RigidSection returns layout function for a form heading
func RigidSection(gtx *layout.Context, th *giomat.Theme, caption string) layout.FlexChild {
    inset := layout.UniformInset(unit.Dp(3))
    return layout.Rigid(func() {
        inset.Layout(gtx, func() {
            giomat.H5(th, caption).Layout(gtx)
        })
    })
}

// RigidLabel returns layout function for a regular label
func RigidLabel(gtx *layout.Context, th *giomat.Theme, caption string) layout.FlexChild {
    return layout.Rigid(func() {
        giomat.Label(th, unit.Px(16), caption).Layout(gtx)
    })
}

// RigidCheckBox returns layout function for a regular checkbox
func RigidCheckBox(gtx *layout.Context, th *giomat.Theme, label string, checkbox *widget.Bool) layout.FlexChild {
    return layout.Rigid(func() {
            giomat.CheckBox(th, label).Layout(gtx, checkbox)
    })
}
