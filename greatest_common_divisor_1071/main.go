package gcd1071

import "fmt"

func gcdOfStrings(w1, w2 string) string {
	if w1+w2 != w2+w1 {
		return ""
	}
	sub := gcd(len(w1), len(w2))
	fmt.Println("sub", sub, "result")
	return string(w1[0:sub])
}

func gcd(l1, l2 int) int {
	if l2 == 0 {
		return l1
	}
	return gcd(l2, l1%l2)
}
