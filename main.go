package main

import (
	"fmt"
)


type Byrunes []rune


func main(){
	fmt.Println("Hello world")
	fmt.Println(IsPalindrome("abba"))
}

func IsPalindrome(text string) bool {
	runes := []rune(text)
	length := len(runes)
	for i, j := 0, length-1; i < length/2; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}



