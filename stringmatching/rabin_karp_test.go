package stringmatching

import (
	"fmt"
	"testing"
)

func TestRabinKarp(t *testing.T) {
	fmt.Println(hash("hi"))
	fmt.Println(hash("abr"))
	fmt.Println(hash("bra"))
	fmt.Println(hashShorten(hash("abr"), 'a'))
	fmt.Println(hash("br"))
	fmt.Println(hashLengthen(hash("br"), 'a'))

	fmt.Println(RabinKarp("abracadabra", "aca"))
	fmt.Println(RabinKarpMulti("abracadabra", 3, "aca", "rac"))
}
