package main

import "fmt"

func isSubsequence(s, t string) bool {
	r := 0
	l := 0
	for l < len(s) && r < len(t) {
		if s[l] == t[r] {
			l++
		}
		r++

	}

	return !(l < len(s))
}

func main() {
	if isSubsequence("abc", "ahgbdc") {
		fmt.Println("sim")
	} else {
		fmt.Println("nao")
	}

}
