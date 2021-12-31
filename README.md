# Gap Buffer Implementation in Go

![Gap Buffer Example](example/example.gif)

## Overview
This package provides a [Gap Buffer](https://en.wikipedia.org/wiki/Gap_buffer) data structure implementation for Go.

Gap Buffer data structure allows efficient insertion and deletion operations clustered near the same location. Gap buffers are especially common in text editors, where most changes to the text occur at or near the current location of the cursor. The text is stored in a large buffer in two contiguous segments, with a gap between them for inserting new text. Moving the cursor involves copying text from one side of the gap to the other (sometimes copying is delayed until the next operation that changes the text). Insertion adds new text at the end of the first segment; deletion deletes it.

## Install Package
```bash
go get github.com/neelanjan00/gap-buffer
```

## Import Package
```go
import gapBuffer "github.com/neelanjan00/gap-buffer"
```

## Operations
The gap buffer implementation provides the following operations:

### 1. SetString
SetString returns the text stored in the buffer.

### 2. GetString
GetString returns the text stored in the buffer.

### 3. MoveCursorRight
MoveCursorRight moves the cursor position to the right by one step.

### 4. MoveCursorLeft
MoveCursorLeft moves the cursor position to the left by one step.

### 5. Delete
Delete deletes a character immediately after the cursor.

### 6. Backspace
Backspace deletes a character immediately before the cursor.

### 7. Insert
Insert inserts a single character at the cursor position.

## Example
```go
package main

import (
	"fmt"

	gapBuffer "github.com/neelanjan00/gap-buffer"
)

func main() {

	// create a GapBuffer struct variable
	gb := new(gapBuffer.GapBuffer)

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
```