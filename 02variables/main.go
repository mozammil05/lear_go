package main

import "fmt"

func main()  {
	var username string = "hello"
	fmt.Println(username);
	fmt.Printf("variable type is: %T  \n", username)

	var isLoaging bool = true
	fmt.Println(isLoaging);
	fmt.Printf("variable type is: %T  \n", isLoaging)

	var smallValue uint8 = 255
	fmt.Println(smallValue);
	fmt.Printf("variable type is: %T  \n", smallValue)

}
