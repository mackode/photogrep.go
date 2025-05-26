package main

import (
	"bufio"
	"flag"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	var fileNames []string
	if flag.NArg() > 0 {
		fileNames = flag.Args()
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fileNames = append(fileNames, strings.TrimSpace(scanner.Text()))
		}
	}

	app := app.New()
	w := app.NewWindow("PhotoGrep")
	cs := NewCsheet(app)
	grid := cs.MakeGrid(fileNames)
	grid.SetMinSize(fyne.NewSize(800, 600))

	submitButton := widget.NewButton("Submit", func() {
		for fileName := range cs.selected {
			fmt.Println(fileName)
		}
		w.Close()
	})

	content := container.NewVBox(grid, submitButton)
	w.SetContent(content)
	w.ShowAndRun()
}