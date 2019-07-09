package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	listAnagrams("helloworld")
}

func contains (s []string, str string) (b bool){
	for item := range s {
		if strings.EqualFold(s[item], str) {
			return true
		}
	}
	return false
}

func factorial(a int) (x int) {
	if a <= 0 {
		return 0
	} else {
		var fact int = 1
		for i := 1; i <= a; i++ {
			fact *= i
		}
		return fact
	}
}

func stringToCharArray(s string) (c []string, e error) {
	s = strings.ToLower(s)
	var stringz []string = strings.Split(s, "")
	var charArray []string 
	for s := range stringz {
		if strings.EqualFold(stringz[s], " ") {
			return nil, errors.New("Word should not contain space.")
		} else {
			charArray = append(charArray, stringz[s])
		}
	}
	return charArray, nil
}

func listAnagrams(x string) {
	s, e := stringToCharArray(x)
	if (e != nil) {
		fmt.Print(e)
		return
	}
	numberOfLetters := len(s)
	repeatedLetters := []string {""}
	for i := 0; i < len(s); i++ {
		if strings.Count(x, s[i]) > 1 {
			if contains(repeatedLetters, s[i]) {
				continue
			}
			repeatedLetters = append(repeatedLetters, s[i])
		}
	}
	if numberOfLetters >= 8 {
		numberOfLetters = 7
	}

	anagrams := []string {x}
	for {
		r := rand.Intn(len(s))
		for i := 0; i < r; i++ {
			x := rand.Intn(len(s))
			s[i], s[x] = s[x], s[i]
		}
		var str string
		for st := range s {
			str = str + s[st]
		}
		var gram string
		var isDuplicate bool = false
		for a := range anagrams {
			gram = anagrams[a]
			if (strings.EqualFold(str, gram)) {
				isDuplicate = true
				break
			}
		}
		if (!isDuplicate) {
			anagrams = append(anagrams, str)
		}
		if len(anagrams) == factorial(numberOfLetters) {
			break
		}
	}

	sort.Strings(anagrams)

	fmt.Println("Generated from", x + ":")
	for anagram := range anagrams {
		fmt.Println(anagrams[anagram])
	}
}
