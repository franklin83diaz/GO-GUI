package ExtryExample

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/widget"
)

func entryNumeric() *widget.Entry {
	entryNumeric := widget.NewEntry()

	entryNumeric.SetPlaceHolder("Enter a number")

	entryNumeric.Validator = func(s string) error {
		// If the string is empty, it is valid
		if s == "" {
			return nil
		}
		// convert the string to a number
		if _, err := strconv.Atoi(s); err != nil {
			// If the string is not a number, return an error
			return fmt.Errorf("'%s' is not a number", s)
		}
		return nil
	}

	return entryNumeric
}
