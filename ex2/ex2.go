package main

import (
	"errors"
	"fmt"
	"log"
)

func recoveryFunction1() {
	v := recover()
	err, ok := v.(error)
	if !ok {
		fmt.Println("Ошибка")
	}
	if err != nil {
		err = errors.New("jjjjj") //errors.wrap у меня не работает, исчезает и я не совсем поняла зачем городить здесь проверку на ошибку она мне кажется избыточной
		log.Printf("recovered %s", v)
		fmt.Println("Ошибка", err)
	}
}

func divis1(a int) int {
	defer recoveryFunction1()
	a = 6 / a // может надо было сделать здесь проверку?
	return a
}

func main() {
	var a int
	a = 0
	fmt.Println(divis1(a))
}
