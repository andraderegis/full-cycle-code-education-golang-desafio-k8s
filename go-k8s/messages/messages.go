package messages

import (
	"fmt"
)

func Greeting(message string) string {
	return fmt.Sprint("<b>", message, "</b>")
}
