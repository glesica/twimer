package main

import (
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

	persister := NewPersister(formatDuration)

	timeDisplay := &widget.Label{
		Text: formatDuration(persister.Elapsed()),
		TextStyle: fyne.TextStyle{
			Bold: true,
		},
	}

	stopwatch := NewStopwatch(func(elapsed, target time.Duration) {
		persister.SetElapsed(elapsed)
		content := formatDuration(elapsed)
		timeDisplay.SetText(content)
	})
	stopwatch.SetElapsed(persister.Elapsed())

	startStopButton := &widget.Button{
		Text:  "Start / Pause",
		Style: widget.PrimaryButton,
		OnTapped: func() {
			stopwatch.Toggle()
		},
	}

	resetButton := &widget.Button{
		Text: "Reset",
		OnTapped: func() {
			stopwatch.Reset()
		},
	}

	copyPathButton := &widget.Button{
		Text: "Copy",
		OnTapped: func() {
			w.Clipboard().SetContent(persister.ElapsedPath())
		},
	}

	w = a.NewWindow("Twimer")
	w.SetPadded(true)
	w.SetFixedSize(true)
	w.Resize(fyne.Size{
		Height: -1,
		Width:  200,
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
