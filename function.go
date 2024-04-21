package main 

//function 
// a code block  that will be used repeatedly  in a program
//a function work when it work

//making a function by using func keyword

/*
func functionName (parameters){

}
*/

//function can take zero or more arguments
//Difference between parameters and arguments
//when you create a function the variable that will be used in function body is parameter
// arguments is the value that will need to give when function is call

/*
func functionName (x int ,y int){

}
*/

// in GO declare type behind the variable name
// two or more consecutive named function parameters share a type
// u can declare behind the all variable

/*
func functionName (x,y int){

}
*/

//Return Function
/*func FunctionName(param1 type, param2 type) return type {
	// code to be executed
	return output
  }
*/

//multiple return function
/*
func functionName (parameter )(x, y return) {
	
	return x,y
}
*/

//Named Return Function 
/*func myFunction(x int, y int) (result int) {
  result = x + y
  return result
}*/

//Naked Return can do in GO but only recommended in short function
/*func myFunction(x int, y int) (result int) {
  result = x + y
  return 
}*/

//Storing  the return value


// func myFunction(x int, y string) (result int, txt1 string) {
// 	result = x + x
// 	txt1 = y + " World!"
// 	return
//   }
  
//   func main() {
// 	a, b := myFunction(5, "Hello")
// 	fmt.Println(a, b)
//   }

// if u don't want to use some return value u can use _

/* func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
  }
  
  func main() {
	 _, b := myFunction(5, "Hello")
	fmt.Println(b)
  } */


//Anonymous function

// using function as a value 
// function can be used as a value to pass the parameters
// these type of function define without name
// also call as lambda

// add := func(a, b int) int {
// 	return a + b
// }

// Calling the anonymous function through the variable.
// result := add(3, 4)
// fmt.Println(result)

//function Closure
// A closure is a function bundled together with its referencing environment,
// which consists of the variables it references that are outside of its local scope.
// In simpler terms, a closure allows a function to access variables from outside its body,
// even when it is executed outside the scope where it was defined.

//In Go, closures are created by defining a function literal (anonymous function) that references variables 
//from its surrounding lexical scope. 
//These variables can then be accessed and modified by the closure even after 
//the defining function has finished execution.

//Example

/*

func main() {
	Function that returns a closure
	counter := func() func() int {
		count := 0 // Local variable captured by the closure
		return func() int {
			count++ // Accessing and modifying the captured variable
			return count
		}
	}

	Creating an instance of the closure
	increment := counter()

	Each time increment is called, it accesses and increments
	the captured variable 'count'
	fmt.Println(increment()) // Output: 1
	fmt.Println(increment()) // Output: 2
	fmt.Println(increment()) // Output: 3
}

*/