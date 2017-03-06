package main

import (
	"bufio"
	c "comm"
	"fmt"
	"os"
)

func main() {
	conn := c.Client()
	conn.c.CliConnect()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
