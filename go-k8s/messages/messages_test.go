package messages

import (
	"testing"
)

func TestExpectGreeting(t *testing.T) {
	const grettingExpected string = "<b>Hello!</b>"

	greeting := Greeting("Hello!")

	if greeting != grettingExpected {
		t.Errorf("Expected greeting is %s, but result is %s", grettingExpected, greeting)
	}

}

func TestUnexpectedGreeting(t *testing.T) {
	const grettingExpected string = "<b>Hello!</b>"

	greeting := Greeting(grettingExpected)

	if greeting == grettingExpected {
		t.Errorf("Expected greeting is %s, but result is %s", grettingExpected, greeting)
	}
}
