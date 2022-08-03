package main

import (
	"fmt"
	"github.com/tjlcast/biu"
)

func main() {
	bs := []byte{1, 2, 3}
	s := biu.BytesToBinaryString(bs)
	fmt.Println(s) //[00000001 00000010 00000011]
	fmt.Println(biu.ByteToBinaryString(byte(3))) //00000011
}