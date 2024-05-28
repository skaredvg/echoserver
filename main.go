package main

import (
	"fmt"
	"log"
)

func main() {
	n := ""
	fmt.Print("Введите данные: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели данные: %s\n", n)
}
