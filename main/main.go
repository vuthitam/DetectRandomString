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

	CheckIsNotRanDom(slice)
}
func CheckIsNotRanDom(s []string) {
	for i := 0; i < len(s); i++ {
		substring := s[i]
		NGram(substring) // n-gram string thành các substring độ dài 3-len(string[i])
	}

}

func NGram(s string) {

}