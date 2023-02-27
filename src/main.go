package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world 2", os.Getenv("ENV"))
}
