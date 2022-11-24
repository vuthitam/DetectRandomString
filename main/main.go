package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Printf("Enter the domain string: ")
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(s)
	slice := strings.Split(s, ".")
	fmt.Println(slice)
	for i := 0; i < len(slice); i++ {
		fmt.Printf(slice[i] + "\n")
	}

	//CheckIsNotRanDom(slice)

}

func CheckIsNotRanDom(s []string) {
	for i := 0; i < len(s); i++ {
		substring := s[i]

		for num := 1; num <= len(substring); num++ {
			n_words := NGram(substring, num)
			for _, word := range n_words{
				if CheckInCorpus(word)
			}
			// n-gram string thành các substring độ dài num-len(string[i])
		}

	}
}
func CheckInCorpus(corpus map[string]int ,s string) bool{
	_, ok:= corpus[s]
	return ok
}
func NGram(s string, num int) []string {
	var array []string
	slice := strings.Split(s, "")

	for i := 0; i < len(s)-num; i++ {
		newslice := slice[i : i+num]
		ngram := strings.Join(newslice, "")
		array = append(array, ngram)
	}
	return array
}
