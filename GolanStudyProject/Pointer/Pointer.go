//Everything is about pass by value
//

//Pointer serves only 1 purpose : sharing
//Pointer shares value across the program boundary.
//There are several types of program boundary.
//The most common one is between function calls
//We can also have a boundary between Goroutine when we will discuss it later.

//When this program starts up, the runtime creates a Goroutine.
//Every Goroutine is separate path of execution that contains instructions that
//contains instructions that needed to be executed by the machine.
//Can also think of Goroutine as a lightweight thread.
//Think program has only 1 Goroutine: The main Goroutine.

//Every Goroutine is given a block of memory, called the 'stack'.
//The stack memory is Go starts out at 2K. It is very small.It can change over time. Dynamic.
//Every time a function is called, a piece of stack is used to help that function run.
// The growing direction of the stack is downward.

//Every function is given a stack frame, memory execution a function
//The size of every stack frame is known at compiler time. No value can be placed on stack
//unless the compiler knows its size ahead of time.
//If we dont know its size of something at compiled time, it has to be on the heap.

//Zero value enables us to initialize every stack frame that we take.
//Stacks are self cleaning. We clean our stack on the way down.
//Every time we make a function, zero value initialization cleaning stack frame.
//We leave that memory on the way up because we dont know if we would need that again.


package main

type user struct {
	name string
	email string
}

func main() {
	//pass by value

	//Declare variable of type int with a value of 10.
	//This value is put on a stack with a value of 10.

	count := 10

	//To get the  address of this value, we use &.
	println("count:\t Value Of[",count,"]Addr Of[",&count,"]")
	//Output:
	//count:	 Value Of[ 10 ]Addr Of[ 0xc000037f70 ]

	//Pass the "Value of" count
	increment1(count)

	// Printing out the result of count. Nothing has changed.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")


	//Pass the "Address of" count
	increment2(&count)

	// Printing out the result of count. count is updated.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// ---------------
	// Escape analysis
	// ---------------

	stayOnStack()
	escapeToHeap()
}

func increment1(inc int){
	//Increment the "Value of" inc.
	inc++
	println("inc1:\tValue Of[",inc,"],Addr Of[",&inc,"]")
}

func increment2(inc *int){
	//increment the "Address of" inc.
	//The * is an operator. It tell us the value of the pointer points to
	*inc++
	println("inc1:\tValue Of[",inc,"],Addr Of[",&inc,"]\tValue Points To[",*inc,"]")
}


//stayOnStack shows how the variable dose not escape.
//Since we know the size of the user value at compiled time,
//the compiler will put this on a stack frame
func stayOnStack() user{
	//In the stayOnStack stack frame, create a value and initialize it.
	u := user{
		name:  "Wyatt",
		email: "wyatt@gmail.com",
	}
	//Take the value and return it, pass back up to main stack frame.
	return u
}

//escapeToHeap shows how the variable escape.
//This looks almost identical to the stayOnStack function.
//It creates a value of type user a d initialize it. It seems like we are doing the same here.
//However, there is one subtle difference:
//We do noy return the value itself but the address of 'u'
//That is the value that is being passed back up the call stack.
//We are using pointer semantic

//You might think about what we have after this call is:
//Main has a pointer to a value that is on stack frame below.
//If this is the case, then we are in trouble.
//Once we come back up the call stack, this memory is there but its is reusable again.
//It is no longer valid.
//Anytime now main makes a function call, we need to allocate the frame and initialize it

//Think about zero value for a second here. It is enable to us to initialize every stack frame that we take.
//Stack are self cleaning. We clean our stack on the way down. Every time we make a function call, zero value,
//initialization, we are cleaning those stack frames. We leave that memory on the way up because we dont know
//if we need that again.

//Back to the example, it is bad because it looks like we take the address of user value,pass it back up to the call
//stack and we now have a pointer which is about to get erased. Therefore, it is not what will happen.

//What actually gonna happen is the idea of escape analysis.
//Because of line "return &u", this value cannot be put inside the stack frame for this function
//so we have to put it out on the heap
//Escape analysis decide what say on stack and what not.
//In the stayOnStack function, because we are passing the copy of the value itself, it is safe to keep these thing
//on stack. But when we SHARE something above the call stack like this,
//escape analysis said this memory is no longer be valid when we get back to main, we must put it out there on the heap.
//Main will end up having a pointer to the heap.

//In fact, this allocation happens immediately on the heap.
//escapeToHeap is gonna have a pointer to the heap. But 'u' gonna base on value semantic.

func escapeToHeap() *user{
	//In the stayOnStack stack frame, create a value and initialize it.
	u := user{
		name:  "Wyatt",
		email: "wyatt@gmail.com",
	}

	return &u
}

//
//What if we run out of stack space?
//

//What happen next is during that function call, there is a little preamble that ask
//Do we have enough stack space? For this frame?
//If yes then no problem because at compiled time we know the size of every frame.
//If no, we have to bigger frame and these value need to be copy over
//The memory on that stack move. It is a trade off. We have to take the cost of copy.
//Because it dose not happen a lot. The benefit of using less memory any Goroutine outrace the cost

// Because stack can grow, no Goroutine can have a pointer to some other Goroutine stack.
// There would be too much overhead for complier to keep track of every pointer. The latency will
// be insane.
// -> The stack for a Goroutine is only for that Goroutine only. It cannot be shared between
// Goroutine.


//Garbage Collection
//-----------------