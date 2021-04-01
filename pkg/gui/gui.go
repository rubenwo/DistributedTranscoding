package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/rubenwo/DistributedTranscoding/pkg/client"
	"time"
)

func Run(client *client.Client) {
	a := app.New()
	w := a.NewWindow("Hello")

	w.CenterOnScreen()
	w.Resize(fyne.Size{
		Width:  1280,
		Height: 720,
	})

	hello := widget.NewLabel("Hello Fyne!")
	launchBtn := widget.NewButton("Launch", func() {
		prog := dialog.NewProgress("Launching servers...", "Progress:", w)
		msgs := make(chan float64)
		go func() {
			for i := 0; i < 100; i++ {
				msgs <- float64(i) / 100
				time.Sleep(time.Millisecond * 100)
			}
			close(msgs)
		}()
		go func() {

			for msg := range msgs {
				fmt.Println(msg)
				prog.SetValue(msg)
			}
			prog.SetValue(1)
			prog.Hide()
		}()
		prog.Show()
	})

	fileButton := widget.NewButton("Pick File", func() {
		dialog.ShowFileOpen(func(r fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			fmt.Println(r.URI())
			if err := client.AddJob(r.URI().Path()); err != nil {
				dialog.ShowError(err, w)
				return
			}
		}, w)
	})

	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
		launchBtn,
		fileButton,
	))

	w.ShowAndRun()
}
