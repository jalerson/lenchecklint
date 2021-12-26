package p

import (
	"fmt"
)

func testCases() {
	id := "id"
	fmt.Println(len(id))

	uuid := "uuid"
	fmt.Println(len(uuid))

	fmt.Println(len("string"))

	text := "text"
	fmt.Println(len(text)) // want "length check of argument text using builtin len function, consider using utf8.RuneCountInString instead"

	Text := "Text"
	fmt.Println(len(Text)) // want "length check of argument Text using builtin len function, consider using utf8.RuneCountInString instead"

	message := "message"
	fmt.Println(len(message)) // want "length check of argument message using builtin len function, consider using utf8.RuneCountInString instead"

	MESSAGE := "MESSAGE"
	fmt.Println(len(MESSAGE)) // want "length check of argument MESSAGE using builtin len function, consider using utf8.RuneCountInString instead"

	msg := "msg"
	fmt.Println(len(msg)) // want "length check of argument msg using builtin len function, consider using utf8.RuneCountInString instead"

	m := struct {
		Text string
	}{
		Text: "text",
	}

	fmt.Println(len(m.Text)) // want "length check of argument Text using builtin len function, consider using utf8.RuneCountInString instead"
}
