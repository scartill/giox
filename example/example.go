package main

import (
	"strconv"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	l "gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/scartill/giox"
	xmat "github.com/scartill/giox/material"
)

var (
	editor              widget.Editor
	checkbox            widget.Bool
	combo               giox.Combo
	comboSelectButton   widget.Clickable
	comboUnselectButton widget.Clickable
)

func main() {
	combo = giox.MakeCombo(
		[]string{
			"Option A",
			"Option B",
		},
		"select an option")

	run := func() {
		w := app.NewWindow()
		loop(w)
	}

	go run()
	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())

	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := l.NewContext(&ops, e)
			mainWindow(gtx, th)
			e.Frame(gtx.Ops)
		}
	}
}

func mainWindow(gtx l.Context, th *material.Theme) {

	for comboSelectButton.Clicked() {
		combo.SelectItem("Option B")
	}

	for comboUnselectButton.Clicked() {
		combo.Unselect()
	}

	children := []l.FlexChild{
		xmat.RigidSection(th, "giox Example"),
		xmat.RigidEditor(th, "Editor example", "<Insert some text here>", &editor),
		xmat.RigidCheckBox(th, "Checkbox example", &checkbox),
		l.Rigid(func(gtx l.Context) l.Dimensions {
			return xmat.Combo(th, &combo).Layout(gtx)
		}),
		xmat.RigidButton(th, "Force select Option B", &comboSelectButton),
		xmat.RigidButton(th, "Unselect", &comboUnselectButton),
		xmat.RigidLabel(th, strconv.FormatBool(checkbox.Value)),
	}

	if combo.HasSelected() {
		children = append(children, xmat.RigidLabel(th, combo.SelectedText()))
	}

	l.W.Layout(gtx, func(gtx l.Context) l.Dimensions {
		return l.Flex{Axis: l.Vertical}.Layout(gtx, children...)
	})
}
