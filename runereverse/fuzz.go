package main

import "fmt"

func Reverse(s string) string {
	r := []rune(s) // HL
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	// START OMIT
	fmt.Println(Reverse("Hello"))
	// END OMIT
}
