package stringmatching

import (
	"fmt"
	"testing"
)

func Test_getNext(t *testing.T) {
	fmt.Println(getNext("abacab"))    // [-1 0 0 1 0 1]
	fmt.Println(getNewNext("abacab")) // [-1 0 -1 1 -1 0]
}

func TestKMP(t *testing.T) {
	source := "abacaabacababb"
	target := "abacab"
	res := KMP(source, target)
	if target != source[res:res+len(target)] {
		t.Errorf("target: %s; res: %s", target, source[res:res+len(target)])
	}
}
