package main

import "fmt"

func isSubsequence(s, t string) bool {
	r := 0
	l := 0
	for {
		if l >= len(s) {
			return true
		}
		if r >= len(t) {
			return false
		}

		if s[l] == t[r] {
			l++
		}
		r++

	}
}

func main() {
	if isSubsequence("abc", "ahgbdc") {
		fmt.Println("sim")
	} else {
		fmt.Println("nao")
	}

}
