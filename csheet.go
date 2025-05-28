package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	con "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/disintegration/imaging"
)

const ThumbSize = float32(200)
const ViewSize = float32(800)

type csheet struct {
	selected map[string]bool
	app      fyne.App
}

func NewCsheet(app fyne.App) *csheet {
	return &csheet{
		selected: map[string]bool,
		app:      app,
	}
}

func (cs *csheet) MakeGrid(fileNames []string) *con.Scroll {
	grid := con.NewGridWithColumns(3)
	scroll := con.NewScroll(grid)

	go func() {
		for _, fileName := range fileNames {
			pick := cs.newPick(fileName)
			grid.Add(pick)
		}
	}()

	return scroll
}

func (cs *csheet) newPick(fileName string) *fyne.Container {
	img, err := imaging.Open(fileName, imaging.AutoOrientation(true))
	if err != nil {
		panic(err)
	}
	thumnail := imaging.Thumnail(img, int(ThumbSize), int(ThumbSize), imaging.Lanczos)

	image := canvas.NewImageFromImage(thumnail)
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(ThumbSize, ThumbSize))

	check := widget.NewCheck("", func(checked bool) {
		if checked {
			cs.selected[fileName] = true
		} else {
			delete(cs.selected, fileName)
		}
	})

	ci := newClickImage(image, func() {
		//toggle
		check.SetChecked(!check.Checked)
	},
		func() {
			fullView := cs.app.NewWindow("Full Image View")
			img, err := imaging.Open(fileName, imaging.AutoOrientation(true))
			if err != nil {
				panic(err)
			}
			fullImage := canvas.NewImageFromImage(img)
			fullImage.FillMode = canvas.ImageFillOriginal
			inspector := NewPan(fullImage)
			fullView.SetContent(inspector.scroll)
			fullView.Resize(fyne.NewSize(ViewSize, ViewSize))
			inspector.Center()
			fullView.Show()
		})

	image.Refresh()
	return con.NewVBox(ci, check)
}