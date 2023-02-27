package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world 1", os.Getenv("ENV"))
}
