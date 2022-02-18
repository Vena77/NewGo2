package main

import (
	"log"
)

func recoveryFunction() {
	v := recover()
	if v != nil {
		log.Printf("recovered %s", v)
	}
}
// Divis Check division zero
func divis(a int) {
	defer recoveryFunction()
	a = 6 / a
}

func main() {
	var a int
	a = 0
	divis(a)
}
