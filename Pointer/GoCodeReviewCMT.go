package main


//Human-written code is held to a higher standard than machine-written code.
func main() {
	//see
	//https://github.com/golang/go/wiki/CodeReviewComments#receiver-type
}
//..........
//Study Notes
//...........

//^Interfaces:

//Go interfaces generally  belong in the packages that uses value of the interface type,
//not the package that implements those value. The implement package should return concrete
//**Pointer & Struct**
//That way new methods can be added to implementations without requiring extensive refactoring.

//DO NOT define interfaces on the implementor side if API "FOR MOCKING",
//instead, design the API so that it can be tested using the public API of the real implementation.

//DO NOT define interfaces before they are used :
//without a realistic example of usage, it is too difficult to see whether an interface is even necessary,
// let alone what methods it ought to contain.

/*package consumer  // consumer.go

type Thinger interface { Thing() bool }

func Foo(t Thinger) string { … }

------------------------------------

package consumer // consumer_test.go
type fakeThinger struct{ … }
func (t fakeThinger) Thing() bool { … }
…
if Foo(fakeThinger{…}) == "x" { … }*/

//----------------------------------

/*// DO NOT DO IT!!!
package producer

type Thinger interface { Thing() bool }

type defaultThinger struct{ … }
func (t defaultThinger) Thing() bool { … }

func NewThinger() Thinger { return defaultThinger{ … } }*/

//Instead return a concrete type and let the consumer mock the producer implementation.

/*package producer

type Thinger struct{ … }
func (t Thinger) Thing() bool { … }

func NewThinger() Thinger { return Thinger{ … } }*/