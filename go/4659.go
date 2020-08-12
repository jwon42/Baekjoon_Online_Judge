package main

import (
	"fmt"
	"strings"
)

func msgPrint(pw string, acceptable bool) {
	if acceptable {
		fmt.Printf("<%s> is acceptable.\n", pw)
	} else {
		fmt.Printf("<%s> is not acceptable.\n", pw)
	}
}

func isVowel(c string) bool {
	if strings.ContainsAny("aeiou", c) {
		return true
	}
	return false
}

func isSame(c1, c2 byte) bool {
	if c1 == c2 {
		return true
	}
	return false
}

func checkPassword(pw string) {
	var vowelCnt, overlapCnt, sameCnt int

	for i := 0; i < len(pw); i++ {
		if i < (len(pw) - 1) {
			if pw[i] != 'e' && pw[i] != 'o' && isSame(pw[i], pw[i+1]) {
				sameCnt++
			}
		}
		if isVowel(string(pw[i])) {
			vowelCnt++
			if i > 0 && i < (len(pw)-1) {
				if isVowel(string(pw[i-1])) && isVowel(string(pw[i+1])) {
					overlapCnt++
				}
			}
		} else {
			if i > 0 && i < (len(pw)-1) {
				if !isVowel(string(pw[i-1])) && !isVowel(string(pw[i+1])) {
					overlapCnt++
				}
			}
		}

	}
	if vowelCnt > 0 && overlapCnt == 0 && sameCnt == 0 {
		msgPrint(pw, true)
	} else {
		msgPrint(pw, false)
	}
}

func main() {
	var pw string

	for pw != "end" {
		fmt.Scanln(&pw)
		if pw == "end" {
			return
		}
		checkPassword(pw)
	}
}
