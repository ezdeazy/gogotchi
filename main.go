package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func echo(r chan string) {
	sb := strings.Builder{}
	for {
		select {
		case v, ok := <-r:
			sb.WriteString(strings.TrimSpace(v))
			sb.WriteString(" lastname")
			fmt.Println(sb.String())
			fmt.Println(ok)
			fmt.Print("Enter your name: ")
			sb.Reset()
		}
	}
}

func main() {
	r := make(chan string, 0)
	go echo(r)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	for {
		name, _ := reader.ReadString('\n')
		r <- name
	}
}
