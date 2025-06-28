package main

import (
	"fmt"
	"os"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("It broke")
	}
	fmt.Printf("%s", home)
}
