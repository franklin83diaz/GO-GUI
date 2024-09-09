package ExtryExample

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// CustomEntry is a widget.Entry with added key event handling
type CustomEntry struct {
	widget.Entry
}

func (c *CustomEntry) TypedKey(key *fyne.KeyEvent) {
	// Call the original key handling
	c.Entry.TypedKey(key)

	// Check if shift was pressed
	if key.Name == fyne.KeyF1 {
		fmt.Println("Show help")
	}
}

func newCustomEntryHelpF1() *CustomEntry {
	entry := &CustomEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}
