package main

//A struct is a collection of fields.

//While arrays are used to store multiple values of
// the same data type into a single variable,
// structs are used to store multiple values
// of different data types into a single variable.

//Syntax
/*
type struct_name struct {
  member1 datatype;
  member2 datatype;
  member3 datatype;
  ...
}

type Person struct {
  name string
  age int
  job string
  salary int
}
*/
//  var person_1 Person
/*
  person_1.name = "Heller"
  person_1.age = 45
  person_1.job = "Teacher"
  person_1.salary = 6000
*/
//or
/*
person_1 := Person {
	fields...
}
*/


/*
Pointers to structs
Struct fields can be accessed through a struct pointer.
person_pointer := &person_1
person_pointer.name
*/

//Pass struct as  Function arguments
/*
func printPerson(person_1 Person) {
  fmt.Println("Name: ", person_1.name)
  fmt.Println("Age: ", person_1.age)
  fmt.Println("Job: ", person_1.job)
  fmt.Println("Salary: ", person_1.salary)
}
*/

