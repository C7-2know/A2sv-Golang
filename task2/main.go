package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	
	fmt.Println(wordCount("hello*")) // prints map[h:1 e:1 l:2 o:1]
	fmt.Println(wordCount("heLLllo*")) // prints map[h:1 e:1 l:4 o:1]
	fmt.Println(wordCount("heLL19llo_90*")) // prints map[0:1 1:1 9:2 e:1 h:1 l:4 o:1]
	fmt.Println(isPalindrome("lkoikl")) // prints false
	fmt.Println(isPalindrome("racecar"))	// prints true

}

func wordCount(str string) map[string]int {
	res := make(map[string]int)
	str= strings.ToLower(str)
	str=regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(str,"")
	for _, ch := range str {
		if _, ok := res[string(ch)]; ok {
			res[string(ch)] += 1

		} else {
			res[string(ch)] = 1
		}
	}
	return res
}

func isPalindrome(word string) bool{
	word=regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(word,"")
	for i:=0;i<=len(word)/2;i++{
		if word[i]!=word[len(word)-1-i]{
			return false
		}
	}
	return true
}
