package main

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/font/gofont"
	
	xmat "github.com/scartill/giox/material"
)

func mainWindow(gtx *layout.Context, th *material.Theme) {

	children := []layout.FlexChild {
		xmat.RigidSection(gtx, th, "giox Examples"),
	}

	layout.W.Layout(gtx, func() {
		layout.Flex{Axis: layout.Horizontal}.Layout(gtx, children...)
	})
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

func main() {
	run := func() {
		w := app.NewWindow()
		loop(w)
	}

	go run()
	app.Main()
}
