package main

// Type parameters in Go allow you to write functions that 
//work with multiple types without sacrificing type safety.


/*
package main

import "fmt"

/Example function with a type parameter
func Print[T any](x T) {
    fmt.Println(x)
}

func main() {
    /Call Print with different types
    Print(42)        
    Print("hello")    
    Print(3.14)       
}

*/

//Comparable

//comparable constraint is used to specify 
//that a type can be compared using the equality operators == and !=. 
//This means that values of a type that satisfies the comparable 
//constraint can be compared for equality or inequality.


//For example, basic types like int, float64, string, bool, etc., are all comparable.
// Struct types are comparable if all their fields are comparable.
// Arrays are comparable if their element types are comparable.

/*
package main

import "fmt"

/Example function with a comparable constraint
func Equals[T comparable](a, b T) bool {
    return a == b
}

func main() {
    fmt.Println(Equals(10, 10))        // Prints: true
    fmt.Println(Equals("hello", "hi")) // Prints: false
    fmt.Println(Equals(3.14, 3.14))    // Prints: true
}

*/

//Generics
//Go also supports generic types.
// A type can be parameterized with a type parameter,
// which could be useful for implementing generic data structures.

/*
func PrintValue[T any](x T) {
    fmt.Println(x)
}
*/
/*
func PrintString[T string](x T) {
    fmt.Println(x)
}
*/




