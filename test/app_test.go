package main

import "testing"

var s1 *string

func instantiate() {
	s2 := new(string)
	*s2 = "andres"

	s1 = s2
}

func TestMain(t *testing.T) {
	instantiate()
	println(*s1)
}
