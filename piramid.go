package main

import (
	"fmt"
	"strconv"
)

func Piramid() {
	var n int
	fmt.Println("############## Piramid ##############")
	fmt.Println("Escribe un número: ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Escriba un entero... ")
	}
	piramidPrim(n)
}

func printH(n int) {
	nStr := strconv.Itoa(n)
	if n <= 0 {
		return
	} else {
		if n == 1 {
			fmt.Println(nStr)
		} else {
			fmt.Print(nStr)
		}
		printH(n - 1)
	}
}

func piramidPrim(n int) {
	if n <= 0 {
		return
	} else {
		printH(n)
		piramidPrim(n - 1)
	}
}
