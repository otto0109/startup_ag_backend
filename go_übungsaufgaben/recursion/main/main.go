package main

import "fmt"

func main() {
	hanoi(8, 0, 1, 2)
}

func hanoi(turmhöhe int, turmposition int, hilfsposition int, turmziel int) {
	if turmhöhe == 0 {
		return
	} else {
		hanoi(turmhöhe-1, turmposition, turmziel, hilfsposition)
		fmt.Println(fmt.Sprintf("Schiebe die Scheibe von %v nach %v\n", turmposition, turmziel))
		hanoi(turmhöhe-1, hilfsposition, turmposition, turmziel)
	}
}

func checkPalindrom(word string) bool {
	if len(word) == 0 {
		return true
	} else {
		return firstLetterIsLastLetter(word) && checkPalindrom(word[1:len(word)-1])
	}
}

func firstLetterIsLastLetter(word string) bool {
	return word[:1] == word[len(word)-1:]
}

func reverse(word string) string {
	if len(word) == 0 {
		return ""
	} else {
		return reverse(word[1:]) + word[:1]
	}
}

func fac(x int) int {
	if x == 0 {
		return 1
	} else {
		return mult(x, fac(x-1))
	}
}

func hoch(x int, y int) int {
	if y == 0 {
		return 1
	} else {
		return mult(x, hoch(x, y-1))
	}
}

func mult(x int, y int) int {
	if x == 0 {
		return 0
	} else {
		return y + mult(x-1, y)
	}
}
