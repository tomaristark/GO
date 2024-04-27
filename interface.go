package main

//In Go, an interface is a type that defines a set of method signatures.
// It doesn't provide the method implementations
// it specifies what methods a type must have

//An interface type is defined as a set of method signatures.
//A value of interface type can hold any value that implements those methods.

/*
/interface
type Shape interface {
    Area() float64 method
}
type Rectangle struct {
    Width  float64
    Height float64
}

/area method for rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
type Circle struct {
    Radius float64
}
/area method for circle
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}
    rectangle := Rectangle{Width: 5, Height: 3}
    circle := Circle{Radius: 4}

	fmt.Printf("Area of rectangle: %.2f\n", rectangle.Area())//15
    fmt.Printf("Area of circle: %.2f\n", circle.Area())//50.24
*/

//User Interface values in Go can be visualized as pairs of a value 
//and its concrete type. They hold a specific underlying value and
// its corresponding type. When you call a method on an interface value,
// Go executes the method associated with its underlying type. 

//If the concrete value inside the interface itself is nil,
// the method will be called with a nil receiver.
//In some languages this would trigger a null pointer exception,
// but in Go it is common to write methods that gracefully handle being called with a nil receiver
//Note that an interface value that holds a nil concrete value is itself non-nil.

//Nil interface values
//A nil interface value holds neither value nor concrete type.
//Calling a method on a nil interface is a run-time error because 
//there is no type inside the interface tuple to indicate which concrete method to call.

//The interface type that specifies zero methods is known as the empty interface:
//interface{}
//An empty interface may hold values of any type. 
//(Every type implements at least zero methods.)
/*
func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
*/

// /An interface can contain the name of one (or more) other interface(s), which is equivalent to
// explicitly enumerating the methods of the embedded interface in the containing interface.
/*
type ReadWrite interface {
		 Read(b Buffer) bool
		 Write(b Buffer) bool
}
type Lock interface {
		 Lock()
		 Unlock()
}
type File interface {
		 ReadWrite
		 Lock
		 Close()
}
*/

//Type assertion
//x.(T)
//x: This is the expression of an interface type. 
//It represents an interface value that you want to assert the underlying concrete type from.
//T: This is the type you are asserting x to be. 
//It represents the concrete type you expect x to hold.

//var i interface{} = 42    // i is an interface{} containing an int value
//v, ok := i.(int)          // Type assertion, asserting i to be of type int
/*
package main

import "fmt"

/Define an interface
type Shape interface {
    Area() float64
}

/Define a concrete type Circle
type Circle struct {
    Radius float64
}

/Implement the Area method for Circle
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    Create a Circle
    var shape Shape = Circle{Radius: 5}

    /Attempt to assert the underlying type of shape to Circle
    if circle, ok := shape.(Circle); ok {
        fmt.Println("Area of Circle:", circle.Area())
    } else {
        fmt.Println("shape is not a Circle")
    }
}
*/

//Type Switch
//A type switch is like a regular switch statement,
// but the cases in a type switch specify types (not values),
// and those values are compared against the type of the value held by the given interface value.
//
//Certainly! A type switch in Go allows you to test whether 
//an interface value holds a value of a specific type and perform
// different actions based on the type. It's similar to a regular switch statement
// but works with types instead of values.

/*
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}

*/
/*
package main

import (
    "fmt"
)

/Define an interface
type Shape interface {
    Area() float64
}

/Define a concrete type Circle
type Circle struct {
    Radius float64
}

/Implement the Area method for Circle
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

/Define a concrete type Rectangle
type Rectangle struct {
    Width  float64
    Height float64
}

/Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    /Create a Circle and a Rectangle
    shapes := []Shape{Circle{Radius: 5}, Rectangle{Width: 4, Height: 3}}

    /Iterate over the shapes and determine their types
    for _, shape := range shapes {
       / Type switch to determine the type of shape
        switch s := shape.(type) {
        case Circle:
            fmt.Printf("Circle with radius %.2f has area %.2f\n", s.Radius, s.Area())
        case Rectangle:
            fmt.Printf("Rectangle with width %.2f and height %.2f has area %.2f\n", s.Width, s.Height, s.Area())
        default:
            fmt.Println("Unknown shape")
        }
    }
}
*/

//Stringer method
//toString

/*
package main

import "fmt"

/Define a custom type Person
type Person struct {
    Name string
    Age  int
}

/Implement the String method for Person
func (p Person) String() string {
    return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func main() {
   / Create a Person instance
    p := Person{Name: "Alice", Age: 30}

   / Convert Person to a string using Stringer interface
    fmt.Println(p)
}

*/

//the Person type implements the Stringer interface by providing 
//a String() method that returns a string representation of 
//the person's name and age. 
//When we print a Person instance using fmt.Println(p),
// Go automatically calls the String() method on the Person value p, 
//providing a custom string representation.

//In Go, errors are represented by values of the error type, 
//which is an interface type defined as:

/*
type error interface {
    Error() string
}
*/

//The error interface has a single method Error() 
//that returns a string describing the error.
//When a function encounters an error condition,
// it typically returns an error value to indicate
// that something went wrong. If the function execution was successful,
//it returns nil to indicate no error.

/*
package main

import (
    "errors"
    "fmt"
)

/Divide divides x by y and returns the result.
/It returns an error if y is zero.
func Divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, errors.New("division by zero")
    }
    return x / y, nil
}

func main() {
    result, err := Divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
*/