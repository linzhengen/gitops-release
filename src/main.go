package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world 5", os.Getenv("ENV"))
}
