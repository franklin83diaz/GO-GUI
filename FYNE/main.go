package main

import (
	"fmt"

	w "fynedemo/widgets/entryExample"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Fyne Demo")

	// Handle key events at the window level
	myWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyF1 {
			fmt.Println("General help")
		}
	})

	myWindow.SetContent(w.EntryAndCustomEntry())
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.ShowAndRun()

}
