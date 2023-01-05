package main

import "fmt"

const NMAX = 10
type tabInt [NMAX]int

func main() {
	var arr tabInt = [NMAX]int{9, 3 ,5, 7, 1, 2, 4, 6, 8, 0}
	printRekursif(arr, 0)
}

func printRekursif(arr tabInt, i int) {
	if i >= NMAX {
		fmt.Println()
		return
	}
	fmt.Print(arr[i], " ")
	printRekursif(arr, i+1)
}

