package main

//Defer 

// In Go, defer is a keyword used to schedule a function call to be executed 
// at the end of the enclosing function, just before it returns.
// The primary purpose of defer is to ensure that resources are released or 
// cleanup actions are performed regardless of how the function exits—whether it returns normally,
// panics, or encounters a runtime error.


/*
func main() {
    defer fmt.Println("world")

    fmt.Print("hello ")
}
*/

//In this example, "world" will be printed after "hello".
//This is because the defer statement schedules the Println("world") call 
//to be executed when the main() function is about to return, after Print("hello ").


// defer statements are executed in Last In, First Out (LIFO) order. 
// This means if there are multiple defer statements in a function,
//  they will be executed in the reverse order they were defined.

/* 
func main() {
    defer fmt.Println("first")
    defer fmt.Println("second")
    defer fmt.Println("third")
    fmt.Println("hello")
}
*/

//Result
/*
hello
third
second
first

*/