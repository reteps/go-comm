package main

import (
	"bufio"
	"fmt"
	c "go-comm/comm"
	"os"
)

func main() {
	conn := c.Client()
	conn.CliConnect()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		conn.CliSend(text)
		message := conn.CliRead()
		fmt.Println(message)
	}
}