package main

import (
	"fmt"

	gapBuffer "github.com/neelanjan00/gap-buffer"
)

func main() {

	// initialise a GapBuffer struct
	gb := new(gapBuffer.GapBuffer)

	// initialise a string to the buffer
	gb.SetString("A Buffer")

	// display the text from buffer
	fmt.Println(gb.GetString())

	// move the cursor to right
	gb.MoveCursorRight()

	// insert character
	gb.Insert(' ')
	gb.Insert('G')
	gb.Insert('a')
	gb.Insert('p')

	fmt.Println(gb.GetString())

	// move cursor to left
	gb.MoveCursorLeft()
	gb.MoveCursorLeft()
	gb.MoveCursorLeft()

	// backspace
	gb.Backspace()

	fmt.Println(gb.GetString())

	gb.MoveCursorLeft()

	// delete
	gb.Delete()

	fmt.Println(gb.GetString())
}
