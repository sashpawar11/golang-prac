package main

import "fmt"

func main() {
	// string is a readonly sequence of bytes - usually UTF-8
	//
	s := "hé🙂!"
	fmt.Printf("%T\n", s)
	fmt.Println(len(s)) // Counts Bytes not characters ( h takes 1, é takes 2, 🙂 takes 4 , ! takes 1)

	//Indexing returns a single byte,
	//
	fmt.Println(s[2])

	// ints cannot be directly type converted to string valeus
	//
	var c int = 255
	var z string = string(c)

	fmt.Println("int string : ", z) // does not become "241" , rather correspodnign asci value
	/// runes,bytes,strings
	var x rune = 'x'
	var a string = string(x)

	var y byte = 'y'
	var b string = string(y)

	fmt.Println(a, b)

}
