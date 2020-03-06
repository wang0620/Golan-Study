package main

import "fmt"

func main() {

	//Building type
	//=============
	//-Type provide the amount of memory that we can allocate?
	//-What dose that memory represent

	//Type can be specific such as int32 or int64
	//For example
	// -uint8 contains a base 10 number using one byte of memory via Unsigned Integer
	// -int32 contains a base 10 number using 4 byte of memory via 32 bit Integer

	//When we declare a type without being very specific, such as unit or int , it gets mapped
	//based on  the architecture we are building  the code against
	//On a 64 bit architecture, in will map to int64. Similarly , on a 32 bit system, it becomes int32

	//The word size is the number of bytes in a word, whic matches our address size.
	//For example, in 64 bit architecture, the word size is 64 bit (8 bytes), address size is 64
	//bit then our integer should be 64 bit/

	//Zero var concept
	//===============
	//Every single value we create must be initialized . If we don't specify it, it will be sent to the zero var.
	//The zero var. The entire allocation of memory, we reset that bit to 0

	//.eg
	var a int // default var is 0
	var b string// default var is "" empty string
	var c float64// default var is 0
	var d bool// default var is false 0
	// var Point nil
	// complex 0i

	//string are a series of unit8 types
	// A string is a two word data structure :
	//-The first word represent a ptr to backing array
	//-The second word represent its length
	//If it is a zero var then --- 1st word is nil 2nd word is 0.
	//Dynamically

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n",b,b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d float64 \t %T [%v]\n", d, d)

	//Using short declaration operator, we can define and initialize at the same time

	aa := 10
	bb := "Wyatt"//1st word points to an array of characters, 2nd word is 5 bytes
	cc := 3.1415926
	dd := true

	fmt.Printf("aa:=\t %T [%v]\n",aa,aa)
	fmt.Printf("bb:=\t %T [%v]\n",bb,bb)
	fmt.Printf("cc:=\t %T [%v]\n",cc,cc)
	fmt.Printf("dd:=\t %T [%v]\n",dd,dd)

	//Conversion vs Casting

	//Golan dose not have casting, but conversion
	//Instead of telling a compiler to pretend to have some more bytes,
	//We have to allocate more memory
	//Specify type and perform a conversion

	bbb:=10
	fmt.Printf("bbb:=\t %T [%v]\n",bbb,bbb)

	aaa:=int32(10)
	fmt.Printf("aaa:=\t %T [%v]\n",aaa,aaa)




}












