package main

import (
	"fmt"
	"os"
)

func decode() 	{
	fmt.Println("DEcode!")
}

func encode() {
	fmt.Println("ENcode!")
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("No arguments, hence nothing done.")
		return
	}

	cmd := os.Args[1]


	if cmd == "encode" {
		encode()
		return
	}

	if cmd == "decode" {
		decode()
		return
	}
}
