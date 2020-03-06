package main

import "fmt"

//example represent a type with different fields

type example struct {
	flag bool
	counter int16
	pi float32
}

func main() {
	//Declare and initialize
	//----------------------

	//Declare a var of type example set its zero value
	//How much memory do we allocate for example?
	//a bool is 1 byte, int16 is 2 bytes , float32 is 4 bytes
	//Putting together, we have 7 bytes. However, the actual answer is 8.
	//That leads us to a new concept of padding and alignment.
	//The padding byte is sitting between the bool and the int16.
	//The reason is because of alignment.

	//The idea of alignment: It is more efficient for this piece of hardware to
	//read the memory on its alignment boundary*.
	//We will take care of alignment boundary issues so the hardware people dnt.

	//Rule 1:
	//Depending on the size a particular value, Go determines the alignment we need.
	//Evey 2bytes var must follow a 2 bytes boundary. Since the bool var is only 1 byte
	//and start at address 0, then the next int16 must start on address 2.
	//The byte at address that get skipped over become a 1 padding.
	//Similarly, if iy is a 4 bytes var then we will have a 3 bytes padding value.

	var e1 example
	fmt.Printf("%+v\n",e1)
	//Output:
	//{flag:false counter:0 pi:0}

	//Rule 2:
	//The largest field represent the padding for the entire struct.
	//We need to minimize the amount of padding as possible.
	//Always lay out the field from highest to smallest.
	//This will push any padding down to the bottom.

	//In this case, the entire struct size has to follow a 8 bytes value because int64 is bytes.
	// type example struct {
	//     counter int64
	//     pi      float32
	//     float   bool
	// }

	//Declare a variable of type example and init using a struct literal
	//Every line must end with comma.

	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.1415926,
	}

	//Display the field values

	fmt.Printf("%+v\n",e2)

	fmt.Println("Flag ",e2.flag)
	fmt.Println("Counter",e2.counter)
	fmt.Println("Pi",e2.pi)

	//Output:
	//{flag:true counter:10 pi:3.1415925}
	//Flag  true
	//Counter 10
	//Pi 3.1415925


//Declare a variable of an anonymous type and init using a struct literal.
e3 := struct {
	flag bool
	counter int16
	pi float32
}{
	flag: true,
	counter: 10,
	pi: 3.1415926,
}
	fmt.Println("Flag", e3.flag)
	fmt.Println("Counter", e3.counter)
	fmt.Println("Pi", e3.pi)


//Name type vs Anonymous type

//If we have two name type identical struct, we cant assign one to another.
//For example, example1 and example2 are identical struct, var example1 e1 || var example2 e2
//ex1 = ex2 is not allowed. We have to explicitly say that
//ex1 = example2(ex2)
//By performing a conversion
//However, if ex is value of identical anonymous struct type(like e3 above), then it is possible to
//assign ex1 = ex
//.eg

var e4 example
e4 =e3

fmt.Println(e4)
}
