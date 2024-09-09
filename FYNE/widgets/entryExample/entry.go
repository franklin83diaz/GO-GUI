package ExtryExample

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func EntryAndCustomEntry() fyne.CanvasObject {

	// Entry widget
	entry := widget.NewEntry()
	// CustomEntry widget
	customEntry := newCustomEntryHelpF1()
	entryNumeric := entryNumeric()

	customEntry.SetPlaceHolder("Enter text, press F1 for Help...")
	entry.SetPlaceHolder("Enter text...")

	// when the text changes
	entry.OnChanged = func(s string) {
		fmt.Println("Current entry text is:", s)
	}

	//when pressing Enter key
	entry.OnSubmitted = func(s string) {
		fmt.Println("Submitted :", s)
	}

	return container.NewVBox(entry, customEntry, entryNumeric)

}
