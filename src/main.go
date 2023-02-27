package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world 8", os.Getenv("ENV"))
}
