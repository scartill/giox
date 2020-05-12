package main

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/font/gofont"
	
	"github.com/scartill/giox"
	xmat "github.com/scartill/giox/material"
)

var (
	editor widget.Editor
	checkbox widget.Bool
	combo giox.Combo
	comboSelectButton widget.Button
	comboUnselectButton widget.Button
)

func main() {
	combo = giox.MakeCombo(
		[]string {
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
	gofont.Register()
	th := material.NewTheme()
	gtx := new(layout.Context)

	for e := range w.Events() {
		if e, ok := e.(system.FrameEvent); ok {
			gtx.Reset(e.Queue, e.Config, e.Size)
			mainWindow(gtx, th)
			e.Frame(gtx.Ops)
		}
	}
	
	return nil
}

func mainWindow(gtx *layout.Context, th *material.Theme) {

	for comboSelectButton.Clicked(gtx) {
		combo.SelectItem("Option B")
	}

	for comboUnselectButton.Clicked(gtx) {
		combo.Unselect()
	}

	children := []layout.FlexChild {
		xmat.RigidSection(gtx, th, "giox Example"),
		xmat.RigidEditor(gtx, th, "Editor example", "<Insert some text here>", &editor),
		xmat.RigidCheckBox(gtx, th, "Checkbox example", &checkbox),
		layout.Rigid(func() {
			xmat.Combo(th).Layout(gtx, &combo)
		}),
		xmat.RigidButton(gtx, th, "Force select Option B", &comboSelectButton),
		xmat.RigidButton(gtx, th, "Unselect", &comboUnselectButton),
	}

	if combo.HasSelected() {
		children = append(children, xmat.RigidLabel(gtx, th, combo.SelectedText()))
	}

	layout.W.Layout(gtx, func() {
		layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
	})
}
