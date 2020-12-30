package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"

	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("Shadow Counter")

	hello2 := NewTappableLabel("hello fyne!", nil)
	hello1 := NewTappableLabel("hello fyne!", func() {
		hello2.SetText("bye fyne!")
	})

	hello3 := widget.NewLabel("hello fyne!")
	hello1.OnTapped = func() {
		hello1.SetText("bye fyne!")
		hello2.SetText("bye fyne!")
		hello3.SetText("bye fyne!")
	}
	hello4 := widget.NewLabel("hello fyne!")

	println(hello1.Size().Height)
	println(hello1.MinSize().Height)

	layout1 := NewGridWrapLayout(fyne.NewSize(hello1.MinSize().Width, hello1.MinSize().Height))
	container := fyne.NewContainerWithLayout(layout1, hello1, hello2, hello3, hello4)
	window.SetContent(container)

	window.Resize(fyne.NewSize((hello1.Size().Width*4+theme.Padding()*4)+10, hello1.Size().Height))
	window.Resize(fyne.NewSize((hello1.Size().Width*4+theme.Padding()*4)+10, hello1.Size().Height))
	// window.Resize(fyne.NewSize)
	window.FixedSize()
	window.ShowAndRun()
}
