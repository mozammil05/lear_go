package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	welcome:= "Welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name")

	input, _ :=reader.ReadString('\n')
	fmt.Println(input,"thanks")

}