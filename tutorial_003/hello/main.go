package main

import (
	"example/message"
	"fmt"
)

func main() {
	message := message.Hello("Oliver")
	fmt.Println(message)
}
