package main

import (
	"os"

	teller "my-study-go/teller"
)

func main() {
	teller.SayHello()
	teller.Say("Hugo")
	teller.SayGoodbye()
	os.Exit(0)
}
