package main

//What is pointers?
//let's learn with simple example

/*
Imagine that somehow you have 4 GB of data in a variable and you need to pass it to a different function. 
Without a pointer, the entire variable is cloned to the scope of the function that is going to use it. 
So, you'll have 8 GB of memory occupied by using this variable twice that, hopefully, 
the second function isn't going to use in a different function again to raise this number even more.
You could use a pointer to pass a very small reference
 to this chunk to the first function so that just the small reference is cloned
  and you can keep your memory usage low.
 
  /reference - Go Design Patterns (Mario Castro Contreras)
*/
//Go has pointers. A pointer holds the memory address of a value.

// declare pointer
// var ptr *int

//pointer type ptr *

// assign the pointer value &

/*
	Declare a variable of type int
    var x int = 10

    Declare a pointer variable
    var ptr *int

    Assign the address of x to ptr
    ptr = &x

*/



/*

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j


*/

/*
    Print the value of x
    fmt.Println("Value of x:", x)

    Print the value of ptr (memory address)
    fmt.Println("Address stored in ptr:", ptr)

    Access the value stored at the memory address pointed by ptr
    fmt.Println("Value pointed by ptr:", *ptr)

*/