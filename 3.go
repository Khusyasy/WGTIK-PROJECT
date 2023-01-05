package main

import "fmt"

func main() {
	var n int
	fmt.Print("n: ")
	fmt.Scan(&n)
	fmt.Println(digitRekursif(n), "digit")
}

func digitRekursif(n int) int {
	if n < 10 {
		return 1
	}
	return 1 + digitRekursif(n / 10)
}

