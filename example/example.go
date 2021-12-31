package main

import (
	"fmt"

	gapbuffer "github.com/neelanjan00/gap-buffer"
)

func main() {

	// create a GapBuffer struct variable
	gb := new(gapbuffer.GapBuffer)

	// initialize a string to the buffer
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
