package main

//Maps are used to store data values in key:value pairs.

//Each element in a map is a key:value pair.

//A map is an unordered and changeable collection 
//that does not allow duplicates.

/*
var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
b := map[KeyType]ValueType{key1:value1, key2:value2,...}
*/


//create map with make
/*
var a = make(map[KeyType]ValueType)
b := make(map[KeyType]ValueType)
*/

// you can't create empty map with make

//invalid key  types in map 
//Slices Maps Functions

//Accessing Map Elements
//value = map_name[key]

//updating map elements
//map_name[key] = value
//adding
//map_name[key]=value

//deleting
//delete(map_name, key)


// checking specific elements in map
//var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day":""}

//val1, ok1 := a["brand"] // Checking for existing key and its value
// val2, ok2 := a["color"] // Checking for non-existing key and its value
//val3, ok3 := a["day"]   // Checking for existing key and its value
//_, ok4 := a["model"]    // Only checking for existing key and not its value

//?Maps are references to hash tables.