package main

//The range form of the for loop iterates over a slice or map.
//When ranging over a slice, two values are returned for each iteration. 
//The first is the index, and the second is a copy of the element at that index.


//You can skip the index or value by assigning to _.
//for i, _ := range array
//for _, value := range array

//If you only want the index, you can omit the second variable.
//for i := range array