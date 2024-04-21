package main

//unlike arrays, the length of a slice can grow and shrink

//Creating a slice
//var slice_name = []datatype{values}
//or
// slice_name := []datatype{values}

//creating a slice from array
//  arr1 := [6]int{10, 11, 12, 13, 14,15}
// my_slice := arr1[2:4] // [12 13]
// A slice is formed by specifying two indices, a low and high bound,
// separated by a colon:
//This selects a half-open range which includes the first element, but excludes the last one.

//? Slices are like references to arrays
//A slice does not store any data, it just describes a section of an underlying array.
//mean changing tha value in slice effect the original or underlying array
/*
names := [4]string{ //array
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)[John Paul George Ringo]

	a := names[0:2][John Paul]
	b := names[1:3] [Paul George]
	fmt.Println(a, b) [John Paul] [Paul George]

	b[0] = "XXX" b[0] is Paul
	fmt.Println(a, b)[John XXX] [XXX George]
	fmt.Println(names) original array [John XXX George Ringo]
*/

//Create slice with make function
//slice_name := make([]type, length, capacity)
//if don't declare capacity it will take the value of length

// The length of a slice is the number of elements it contains
//The capacity of a slice is the maximum number of elements that the slice can hold without allocating more memory
//if u declare capacity is 10 that will take 10 memory without using it	 or not

//my_slice := make([]int, 5, 10)

//The zero value of a slice is nil.

//A nil slice has a length and capacity of 0 and has no underlying array.



//Appending slice


//slice_name = append(slice_name, element1, element2, ...)
//my_slice1 := []int{1, 2, 3, 4, 5, 6}
//my_slice1 = append(my_slice1, 20, 21)


//Slices can contain any type, including other slices.
//appending slice
//slice3 = append(slice1, slice2...)

/*
  my_slice1 := []int{1,2,3}
  my_slice2 := []int{4,5,6}
  my_slice3 := append(my_slice1, my_slice2...)
*/


//When using slices, Go loads all the underlying elements into the memory.
//If the array is large and you need only a few elements, 
//it is better to copy those elements using the copy() function.
//The copy() function creates a new underlying array 
//with only the required elements for the slice. 
//This will reduce the memory used for the program. 

//copy(dest, src)


//numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
// neededNumbers := numbers[:len(numbers)-10]
//numbersCopy := make([]int, len(neededNumbers))
// copy(numbersCopy, neededNumbers)

//The copy() function takes in two slices dest and src, and 
//copies data from src to dest.
// It returns the number of elements copied.