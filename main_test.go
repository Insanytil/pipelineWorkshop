package main

import "testing"

func TestAdd(t *testing.T) {
	res := Add(1, 2)

	if res != 3 {
		t.Error("Add is not working")
	}

   res = Add(2, 3)

   if res != 5 {
		t.Error("Add is not working properly")
   }


}



func TestSub(t *testing.T) {
	res := Sub(5, 2)

	if res != 3 {
		t.Error("Sub is not working")
	}
}
