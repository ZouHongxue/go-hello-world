package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	p := "paech"

	pp := "peach punch"

	ppp := "peach punch pinch"

	match, _ := regexp.MatchString("p([a-z]+)ch", p)
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString(pp))

	fmt.Println(r.FindStringIndex("peach punch"))

	fmt.Println(r.FindStringSubmatch("peach punch"))

	fmt.Println(r.FindStringSubmatchIndex("peach pouch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex(ppp, -1))

	fmt.Println(r.FindAllString(ppp, 2))

	fmt.Println(r.Match([]byte(p)))

	r = regexp.MustCompile("p([a-z]+)ch")

	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")

	out := r.ReplaceAllFunc(in, bytes.ToUpper)

	fmt.Println(string(out))

}
