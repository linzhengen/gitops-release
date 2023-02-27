package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world 6", os.Getenv("ENV"))
}
