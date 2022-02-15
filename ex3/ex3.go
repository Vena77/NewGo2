package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i := 0; i < 1000000; i++ {
		file, err := os.Create("files/hello" + strconv.Itoa(i) + ".txt")

		if err != nil {
			fmt.Println("Unable to create file:", err)
			os.Exit(1)
		}
		defer file.Close()
	}
}
