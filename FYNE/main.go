package main

import (
	"fmt"
	"os"

	w "fynedemo/widgets/entryExample"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {

	if len(os.Args) < 2 {
		invalid()
		return
	}

	//TODO: refactor this code to use a switch statement
	if os.Args[1] == "entry" {

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
	} else {
		invalid()
	}

}

func invalid() {
	fmt.Println("Invalid argument")
	fmt.Println("List of valid arguments: ")
	//set color green
	fmt.Print("\033[32m")
	fmt.Println("entry")
	//reset color
	fmt.Print("\033[0m\n")
}
