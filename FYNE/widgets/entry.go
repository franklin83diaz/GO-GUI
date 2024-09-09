package widgets

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func EntryAndCustomEntry() fyne.CanvasObject {

	entry := widget.NewEntry()
	customEntry := NewCustomEntry()
	entry.SetPlaceHolder("Enter text...")

	// when the text changes
	entry.OnChanged = func(s string) {
		fmt.Println("Current entry text is:", s)
	}

	//when pressing Enter key
	entry.OnSubmitted = func(s string) {
		fmt.Println("Submitted :", s)
	}

	return container.NewVBox(entry, customEntry)

}
