package main

import (
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile("./files/test.txt")
	check(err)
	fmt.Print(string(dat))

	fmt.Printf("%T", dat)
	// dump := []byte{string(dat)}
	err = os.WriteFile("./files/res.txt", dat, 0644)

	// file, err := os.Open("./files/test.txt")
	// check(err)

	// byte1 := make([]byte, 5)
	// n1, err := file.Read(byte1)
	// check(err)
	// fmt.Printf("%d bytes: %s\n", n1, string(byte1[:n1]))

	// file.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
