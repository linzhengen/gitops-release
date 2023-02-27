package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world 4", os.Getenv("ENV"))
}
