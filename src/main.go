package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world 7", os.Getenv("ENV"))
}
