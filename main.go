package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"time"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	var w fyne.Window

	timeDisplay := &widget.Label{
		Text: "...",
		TextStyle: fyne.TextStyle{
			Bold: true,
		},
	}

	stopwatch := NewStopwatch(func (elapsed, target time.Duration) {
		hours := elapsed / time.Hour
		minutes := (elapsed - (hours * time.Hour)) / time.Minute
		seconds := (elapsed - (hours * time.Hour) - (minutes * time.Minute)) / time.Second
		content := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
		timeDisplay.SetText(content)
	})

	startStopButton := &widget.Button{
		Text:  "Start / Pause",
		Style: widget.PrimaryButton,
		OnTapped: func() {
			println("[DEBUG] Start / Pause")
			stopwatch.Toggle()
		},
	}

	resetButton := &widget.Button{
		Text: "Reset",
		OnTapped: func() {
			println("[DEBUG] Reset")
			stopwatch.Reset()
		},
	}

	copyPathButton := &widget.Button{
		Text: "Copy",
		OnTapped: func() {
			w.Clipboard().SetContent("/home/stuff/file.txt")
		},
	}

	w = a.NewWindow("Twimer")
	w.SetPadded(true)
	w.SetFixedSize(true)
	w.Resize(fyne.Size{
		Height: -1,
		Width: 200,
	})
	w.SetContent(&widget.Box{
		Children: []fyne.CanvasObject{
			&fyne.Container{
				Layout: layout.NewGridLayout(2),
				Objects: []fyne.CanvasObject{
					resetButton,
					copyPathButton,
				},
			},
			&fyne.Container{
				Layout: layout.NewCenterLayout(),
				Objects: []fyne.CanvasObject{
					timeDisplay,
				},
			},
			startStopButton,
		},
	})

	w.ShowAndRun()
}
