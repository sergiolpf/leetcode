package main

import (
	"fmt"
	"strconv"
)

func isIsomorphic(s string, t string) bool {
	hashS := make(map[string]int)
	hashT := make(map[string]int)

	numberS := 1
	numberT := 1

	for i := 0; i < len(s); i++ {
		_, isFound := hashS[string(s[i])]
		if !isFound {
			hashS[string(s[i])] = numberS
			numberS += 1
		}

		isFound = false

		_, isFound = hashT[string(t[i])]
		if !isFound {
			hashT[string(t[i])] = numberT
			numberT += 1
		}
	}

	valueS := ""
	valueT := ""
	for i := 0; i < len(s); i++ {
		valueS += strconv.Itoa(hashS[string(s[i])])
		valueT += strconv.Itoa(hashT[string(t[i])])

	}

	return valueS == valueT

}

func main() {
	if isIsomorphic("paper", "title") {
		fmt.Println("true")
	} else {

		fmt.Println("false")
	}
}
