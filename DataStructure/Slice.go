package main
// Reference types: slice, map, channel, interface, function.
// Zero value of a reference type is nil.

import (
	"fmt"
)
func main() {

	//----------------------
	//Declare and initialize
	//----------------------

	//Create a slice with length of 5 elements.
	//make is a special built-in function that only works with slice, map and channel.
	//make creates a slice that has an array of 5 strings behind it.
	// We are getting 3 word data structure

	//The first word points to the backing array
	//Second word is length
	//Third one is capacity

	// |  *  | --> | nil | nil | nil | nil | nil |
	//  -----      |  0  |  0  |  0  |  0  |  0  |
	// |  5  |
	//  -----
	// |  5  |
	//  -----


	//---------------------
	//Length VS  Capacity
	//---------------------

	//Length is the number of elements from pointer position we have access to (read and write).
	//Capacity is total number of elements from this pointer position that exist in the backing array

	//Syntactic sugar ->> looks like array
	//It also have the some cost that we have seen in array.
	//One thing mindful about: there is no value in bracket []string inside the make function.
	//With that in mind, we can constantly notice that we are dealing with a slice, not array

	slice1 := make([]string, 5)

	slice1[0] = "apple"
	slice1[1] = "feojia"
	slice1[2] = "kiwi"
	slice1[3] = "strawberry"
	slice1[4] = "blood orange"

	//We can not access an index of a slice beyond its length
	//Error: panic: runtime error: index out of range
	//slice1[5]="Runtime error"

	//We are passing the value of slice, noy its address. So the Println function will have its own copy of the slice.

	fmt.Println("Printing a slice")
	fmt.Println(slice1)

	//Reference type

	//Create a slice with a length of 5 elements and a capacity of 8.
	//make allows us to adjust the capacity directly on construction of initialization.
	//What we end up having now is a 3 word data structure where the first word points to an array
	//of 8 elements, length is 5 and capacity is 8.

	// |  *  | --> | nil | nil | nil | nil | nil | nil | nil | nil |
	//  -----      |  0  |  0  |  0  |  0  |  0  |  0  |  0  |  0  |
	// |  5  |
	//  -----
	// |  8  |
	//  -----
	// It means that I can read and write to the first 5 elements and I have 3 elements of capacity
	// that I can leverage later.
	//.................
	slice2 := make([]string, 5, 8)
	slice2[0] = "Apple"
	slice2[1] = "Orange"
	slice2[2] = "Banana"
	slice2[3] = "Grape"
	slice2[4] = "Plum"

	fmt.Printf("\n=> Length vs Capacity\n")
	inspectSlice(slice2)

	//============================================================
	//The idea of appending: making slice a dynamic data structure
	//============================================================

	fmt.Println("The Idea of appending")

	//Declare a nil slice of strings, set to its zero value.
	//3 word data structure:
	//First one points to nil,
	//Second and last are zero.

	var data []string
	//inspectSlice(data)

	//What if I do data:= string{}? Is it the same?
	//No because data in this case is not set to its zero value.
	//This is why we always user var for zero value because not every type when we create an empty
	//literal we have its zero value in return.

	//What actually happen here is that we have a slice but it has a pointer (as opposed to nil).
	//This is consider an empty slice, not a nil slice.
	//There is a semantic between a nil and an empty slice. Any reference type that set to iys zero
	//value can be considered nil. If we pass a nil slice to marshal function, we get back a string that said null
	//but when we pass an empty slice, we get an empty JSON documents.
	//But where dose that pointer point to? It is empty struct, which we will review it later.

	//Capture the capacity of the slice.
	lastCap :=cap(data)

	//Append ~100k strings to the slice.
	for record :=1; record <= 102400; record++{
		//Use the built-in function append to add to the slice.
		//It allows us to add value to a slice, making the data structure dynamic, yet still allow us to
		//use that contiguous block of memory that give us the predictable access patten from mechanical sympathy.
		//The appending call is working with value semantic. We are not sharing this slice but appending to it and
		//returning a new copy of it. The slice gets ti stay on the stack, No heap
	}







	
}


// inspectSlice exposes the slice header for review.
// Parameter: again, there is no value in side the []string so we want a slice.
// Range over a slice, just like we did with array.
// While len tells us the length, cap tells us the capacity
// In the output, we can see the addresses are aligning as expected.
func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i := range slice {
		fmt.Printf("[%d] %p %s\n", i, &slice[i], slice[i])
	}
}