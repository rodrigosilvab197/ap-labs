package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	
	if len(os.Args) > 1 { 
        name := os.Args[1:]
		fmt.Printf("Welcome, ")
		fmt.Printf(strings.Join(name, " "))
		fmt.Printf(" to the jungle \n")
    } else {
        fmt.Println("Sorry bro! input no found ! try again!")
    }
	
	
}
