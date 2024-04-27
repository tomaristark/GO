package main

//A method is a function with a special receiver argument. 
//! receiver == argument
//The receiver appears in its own argument list between the func keyword and the method name.
/*
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

*/

//You can declare a method on non-struct types, too.
//!You can only declare a method with a receiver whose type is
//! defined in the same package as the method
//You cannot declare a method with a receiver whose type is defined in another package

//Method with Pointer receiver (argument)
//You can declare methods with pointer receivers.

/*
type Person struct {
	name string
	age	 int
	gender string
}
func(*Person)isMale bool {
	
}
*/
//Methods with pointer receivers can modify the value to which the receiver point.
//methods often need to modify their receiver, 
//pointer receivers are more common than value receivers.


//?Choosing a value or pointer receiver
//There are two reasons to use a pointer receiver.
//The first is so that the method can modify the value that its receiver points to.
//The second is to avoid copying the value on each method call.
// This can be more efficient if the receiver is a large struct, for example.





