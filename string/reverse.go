package main

import "fmt"

func reverse(s []byte, st, e int) {
	if st >= e {
		return
	}
	s[st], s[e] = s[e], s[st]
	st++
	e--
	reverse(s, st, e)
}
func ReverseString(s string) string {
	str := []byte(s)
	reverse(str, 0, len(s)-1)
	return string(str)
}
func main() {
	s := ReverseString("vishnu")
	fmt.Println(s)
}
