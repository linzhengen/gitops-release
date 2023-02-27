package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world 3", os.Getenv("ENV"))
}
