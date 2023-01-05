package main

import "fmt"

func main() {
	var n int
	fmt.Print("n: ")
	fmt.Scan(&n)
	fmt.Println(binerRekursif(n))
}

func binerRekursif(d int) int {
	if d == 1 {
		return 1
	}
	return (binerRekursif(d / 2) * 10) + (d % 2)
}

