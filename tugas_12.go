package main

import (
	"fmt"
	"regexp"
)

func main() {
	es := `(\bcats?\b)|(\bdogs?\b)|(\brats?\b)`
	e := regexp.MustCompile(es)

	fmt.Println("Program Mencocokan Kata dengan Regex")
	fmt.Println("Kalimat mengandung cats, dogs, rats")
	fmt.Println("Masukan Kata pertama:")
	var kalimat1 string
	fmt.Scanln(&kalimat1)
	fmt.Println("Masukan Kata kedua:")
	var kalimat2 string
	fmt.Scanln(&kalimat2)
	fmt.Println("Masukan Kata ketiga:")
	var kalimat3 string
	fmt.Scanln(&kalimat3)

	match(e, kalimat1)
	match(e, kalimat2)
	match(e, kalimat3)
}

func match(r *regexp.Regexp, text string) {
	matched := r.MatchString(text)
	if matched {
		fmt.Println("Kata cocok", r.String(), ":", text)
	} else {
		fmt.Println("Kata tidak cocok", r.String(), ":", text)
	}
}
