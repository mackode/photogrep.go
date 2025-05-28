package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type clickImage struct {
	widget.BaseWidget
	image  *canvas.Image
	cbleft func()
	cbright func()
}

func newClickImage(image *canvas.Image, cbleft, cbright func()) *clickImage {
	ci := &clickImage{}
	ci.ExtendBaseWidget(ci)
	ci.image = img
	ci.cbleft = cbleft
	ci.cbright = cbright
	return ci
}

func (t *clickImage) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(t.image)
}

func (t *clickImage) Tapped(ev *fyne.PointEvent) {
	t.cbleft()
}