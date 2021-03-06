package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"github.com/rubenwo/DistributedTranscoding/pkg/client"
	"image/color"
	"math"
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

	//hello := widget.NewLabel("Hello Fyne!")
	//launchBtn := widget.NewButton("Launch", func() {
	//	prog := dialog.NewProgress("Launching servers...", "Progress:", w)
	//	msgs := make(chan float64)
	//	go func() {
	//		for i := 0; i < 100; i++ {
	//			msgs <- float64(i) / 100
	//			time.Sleep(time.Millisecond * 100)
	//		}
	//		close(msgs)
	//	}()
	//	go func() {
	//
	//		for msg := range msgs {
	//			fmt.Println(msg)
	//			prog.SetValue(msg)
	//		}
	//		prog.SetValue(1)
	//		prog.Hide()
	//	}()
	//	prog.Show()
	//})
	//
	//fileButton := widget.NewButton("Pick File", func() {
	//	dialog.ShowFileOpen(func(r fyne.URIReadCloser, err error) {
	//		if err != nil {
	//			dialog.ShowError(err, w)
	//			return
	//		}
	//		fmt.Println(r.URI())
	//		if err := client.AddJob(r.URI().Path()); err != nil {
	//			dialog.ShowError(err, w)
	//			return
	//		}
	//	}, w)
	//})

	//w.SetContent(container.NewVBox(
	//	hello,
	//	widget.NewButton("Hi!", func() {
	//		hello.SetText("Welcome :)")
	//	}),
	//	launchBtn,
	//	fileButton,
	//))
	//myCanvas := w.Canvas()
	//
	//text := canvas.NewText("Text", color.White)
	//text.TextStyle.Bold = true
	//myCanvas.SetContent(text)
	//go changeContent(myCanvas)

	fractal := &fractal{window: w}
	fractal.canvas = canvas.NewRasterWithPixels(fractal.mandelbrot)

	fractal.currIterations = 100
	fractal.currScale = 1.0
	fractal.currX = -0.75
	fractal.currY = 0.0

	w.SetContent(fyne.NewContainerWithLayout(fractal, fractal.canvas))

	w.ShowAndRun()
}

func changeContent(c fyne.Canvas) {

	time.Sleep(time.Second * 2)

	c.SetContent(canvas.NewRectangle(color.Black))

	time.Sleep(time.Second * 2)
	c.SetContent(canvas.NewLine(color.Gray{0x66}))

	time.Sleep(time.Second * 2)
	raster := canvas.NewRasterWithPixels(func(x, y, w, h int) color.Color {
		return color.RGBA{
			R: 255,
			G: 234,
			B: 0,
			A: 175,
		}
	})

	c.SetContent(raster)
	time.Sleep(time.Second * 2)

	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = color.RGBA{0xff, 0x33, 0x33, 0xff}
	c.SetContent(circle)

	time.Sleep(time.Second * 2)
	c.SetContent(canvas.NewImageFromResource(theme.FyneLogo()))
}

type fractal struct {
	currIterations          uint
	currScale, currX, currY float64

	window fyne.Window
	canvas fyne.CanvasObject
}

func (f *fractal) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	f.canvas.Resize(size)
}

func (f *fractal) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(320, 240)
}

func (f *fractal) refresh() {
	if f.currScale >= 1.0 {
		f.currIterations = 100
	} else {
		f.currIterations = uint(100 * (1 + math.Pow((math.Log10(1/f.currScale)), 1.25)))
	}

	f.window.Canvas().Refresh(f.canvas)
}

func (f *fractal) scaleChannel(c float64, start, end uint32) uint8 {
	if end >= start {
		return (uint8)(c*float64(uint8(end-start))) + uint8(start)
	}

	return (uint8)((1-c)*float64(uint8(start-end))) + uint8(end)
}

func (f *fractal) scaleColor(c float64, start, end color.Color) color.Color {
	r1, g1, b1, _ := start.RGBA()
	r2, g2, b2, _ := end.RGBA()
	return color.RGBA{f.scaleChannel(c, r1, r2), f.scaleChannel(c, g1, g2), f.scaleChannel(c, b1, b2), 0xff}
}

func (f *fractal) mandelbrot(px, py, w, h int) color.Color {
	drawScale := 3.5 * f.currScale
	aspect := (float64(h) / float64(w))
	cRe := ((float64(px)/float64(w))-0.5)*drawScale + f.currX
	cIm := ((float64(py)/float64(w))-(0.5*aspect))*drawScale - f.currY

	var i uint
	var x, y, xsq, ysq float64

	for i = 0; i < f.currIterations && (xsq+ysq <= 4); i++ {
		xNew := float64(xsq-ysq) + cRe
		y = 2*x*y + cIm
		x = xNew

		xsq = x * x
		ysq = y * y
	}

	if i == f.currIterations {
		return theme.BackgroundColor()
	}

	mu := (float64(i) / float64(f.currIterations))
	c := math.Sin((mu / 2) * math.Pi)

	return f.scaleColor(c, theme.PrimaryColor(), theme.TextColor())
}

func (f *fractal) fractalRune(r rune) {
	if r == '+' {
		f.currScale /= 1.1
	} else if r == '-' {
		f.currScale *= 1.1
	}

	f.refresh()
}

func (f *fractal) fractalKey(ev *fyne.KeyEvent) {
	delta := f.currScale * 0.2
	if ev.Name == fyne.KeyUp {
		f.currY -= delta
	} else if ev.Name == fyne.KeyDown {
		f.currY += delta
	} else if ev.Name == fyne.KeyLeft {
		f.currX += delta
	} else if ev.Name == fyne.KeyRight {
		f.currX -= delta
	}

	f.refresh()
}
