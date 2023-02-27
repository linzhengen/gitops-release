package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world 9", os.Getenv("ENV"))
}
