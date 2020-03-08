package main

//--------
//CPU Cache
//---------

//Cores DO NOT access main memory directly but their CACHE
//What store in cache is date an structure

//Cache speed fast to low is : L1 >> L2 >> L3 -> main memory.

//Scott Meyers: "If performance matter then
//total memory you have is the total amount f caches"
//-> Access to main memory is incredibly slow; Practically speaking it might not even be there

//How do we write code that can be sympathetic with the caching system to make sure that
//we do not have a cache miss or least, we minimize cache misses to our fullest potential?

//Processor has a "Pre-fetcher".It predicts what data is needed ahead of time.
//There are different  granularity depending on where we are on the machine.
//Our programming model uses a byte. We can read and write to a byte ar a time.
//However, from the caching system POV, our granularity is not 1 byte. It is 64 bytes,
//called a cache line. All memory us junked up in this 64 bytes cache line.


//Since the caching mechanism is complex, Pre-fetcher tries to hide all the latency from us.
//It has to be able to pick up on predictable access patten to data.
//-> We need to write the code that predictable access pattern to data.


//One easy way is to create contiguous allocation of memory and to iterate over them.
//The 'array' data structure gives us ability to do so.
//From the hardware perspective, array is most significant data structure.
//From go perspective, Slice is, Array is the backing data structure for Slice (like Vector in C++)

//Once we allocate an array, whatever it sized, every element is equal distant from other element.
//As we iterate over that array, we begin to walk cache line by cache line.
//As Pre-fetcher see that access patten, it can pick it up and hide all the latency from us.

//For example, we have a big NxN matrix. We do LinkedList Traverse, Column Traverse and Row Traverse
//and Benchmark against them.

//Unsurprisingly, Row Traverse has the best performance. It looks like random access memory by cache line
//and create a predictable access patten

//Column Traverse dose not walk through the matrix cache line by cache line. It looks like random access memory pattern
//That is why slowest among those.

//However, that dose not  explain why the LinkedList Traverse's performance is in the middle. We
//just think that it might perform as poorly as the Column Traverse.
//This lead us to another cache: TLB - Translation look aside buffer. Its job is to maintain
//operating system page and offset to physical memory is.

//
//Translation look aside buffer
//

//Back to the different granularity, the caching system moves data in an out of the hard ware at 64 bytes at a time.
//However, the operating system manages memory by paging its 4k (traditional page size for an operating system)
//TLB: For every page that we are managing, let's take our virtual memory addresses because that we use
//(Software run virtual addresses, it is sandbox, that is how we use or share physical memory). And map it to the
//right page and offset for that  physical memory.

//A miss on the TLB can be worse then just the cache miss alone.
//The LinkedList is somewhere in between is because the chance of multiple nodes being on the same page is probably
//pretty good. Even though we can get cache misses because cache line are not necessary in the distance that is predictable
//We probably not have so many TLB cache。

//In the Column Traverse, not only we have cache misses, we probably have a TLB cache miss on every access as well.

//Data-Oriented design matters.
//It is not enough to write the most efficient algorithm, how we access our data can have much more lasting effect on
//performance than the algorithm itself。


import "fmt"


func main() {
	//======================
	//Declare and initialize
	//======================


	//Declare an array of five strings that is initialized to its zero value
	//Recap: a string is a 2 word data structure: a pointer and a length
	//Since this array is set to its zero value, every string in this array is also set to zero value，
	//Which means that each string has the first word pointed to nil and second word is 0;

	//
	//  -----------------------------
	// | nil | nil | nil | nil | nil |
	//  -----------------------------
	// |  0  |  0  |  0  |  0  |  0  |
	//  -----------------------------

	var strings [5]string

	//At index 0, a string now has a pointer to backing array of bytes (character in string)
	//and its length is 5.

	//What is the cost?

	//The cost of this assignment is the cost of copying 2 bytes.
	//We have two string values that have pointers to same backing array of bytes.
	//Therefore, the cost of this assignment is just 2 words.

	//
	//  -----         -------------------
	// |  *  |  ---> | A | p | p | l | e | (1)
	//  -----         -------------------
	// |  5  |                  A
	//  -----                   |
	//                          |
	//                          |
	//     ---------------------
	//    |
	//  -----------------------------
	// |  *  | nil | nil | nil | nil |
	//  -----------------------------
	// |  5  |  0  |  0  |  0  |  0  |
	//  -----------------------------

	strings[0] = "Apple"
	strings[1] = "Orange"
	strings[2] = "Kiwi"
	strings[3] = "Peach"
	strings[4] = "FEIJOA"


	//Iterate over the array of strings

	//Using range, not only we can get the index but also a copy of the value in the array.
	//Fruits is now a string value; its scope is within the for statement.
	//In the first iteration, we have the word "Apple", It is a string that has the first word
	//also points to (1) and the second word is 5.
	//So we now have 3 different string  value all sharing the same backing array.

	//What are we passing to Println Function?
	//We are using value semantic here. We are not sharing our string value. Println is getting its own copy
	//its own string value. It means when we get to println call,
	//there are now 4 string values all sharing the same backing array.

	//We do not want to take an address of string.
	//We know the size of a string ahead of time.
	//-> It has the ability to be on the stack
	//->Not creating allocation
	//->Not causing pressure on the GC
	//->The string has been designed to leverage value mechanic, to stay on the stack, out of the way creating garbage
	//->The only thing that has to be on the heap, if anything is the backing array, which is the thing that being shared

	fmt.Printf("\n=> Iterate over array\n")

	for i,f := range strings{
		fmt.Println(i,f)
	}

	//declare an array of 4 integers that is initialized with some values using literal

	num := [6]int{ 0,1,2,3,4,5}
	//Iterate over the array o numbers using traditional style
	fmt.Printf("\n=> Iterate over array using tradition\n")

	for i:=0;i<len(num);i++{
		fmt.Println(i,num[i])
	}
	fmt.Println()
	//Different type of arrays

	//Declare an array of 5 integers that is initialized to its zero value.
	var five [5]int

	//Declare an array of 4 integers that is initialized some values

	var four= [4]int{1,2,3,4}

	fmt.Println("Different type of arrays")
	fmt.Println(five)
	fmt.Println(four)
	//Then

	//When we try to assign four to five like so
	//five = four ,
	//This can not happen because they have different types (size and representation).
	//The compiler says that : "Can not user fou (type [4]int) as type [5]int in assignment"
	//The size of array makes up its type name : [4]int vs [5]int.
	//Just like we have seen with pointer.
	//The * in *int is not an operator but part of the type name.

	//Unsurprisingly, all array size has been know at compiled time.

	//Contiguous memory allocation

	//Declare an array of 6 strings initialized with values.

	six := [6] string{"a","b","c","d","e","f"}

	//Iterator over the array displaying the value and address of each element.
	//By looking at the output of this printf function, we can see that
	//this array is truly a contiguous block of memory.
	//We know a string is 2 word and depending on computer architecture,
	//It will have x byte.
	//The distance between two consecutive IndexAddr is exactly
	//x Byte
	//v is its own variable on the stack and it has the same address every single time.


	fmt.Println("\n=>Contiguous memory allocation")
	for i, v:= range six{
		fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &six[i])
	}


}
