package main

import (
	"os"
	"log"
	"fmt"
	"time"
)

func getDifference() int64 {
	info, err := os.Stat("ago")

	if err != nil {
		log.Fatal(err)
	}

	modifiedTime := info.ModTime()
	now := time.Now()
	difference := now.Unix() - modifiedTime.Unix()

	return difference
}

func main() {
	fmt.Println("Function Difference: ", getDifference())
}
