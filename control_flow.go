package main


//like normal if else

// can use with () or without ()
//if 
/* if condition {
	if condition true this code will work
 }
 if v := math.Pow(x, n); v < lim {
		return v
	}
*/

// if else 
/* 
if condition {

} else {
  if condition is false this code will work
}
*/

//if else if
/*
if condition1 {
   
} else if condition2 {
   
} else {

}

*/


//for 

/* for initializer ; condition; loop counter {

  }*/

/*
func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }
}

*/

//continue skip the value break for break the loop

/*
forever loop
for{

}
*/

//Switch

// A switch statement is a shorter way to write a sequence of if - else statements.
//  It runs the first case whose value is equal to the condition expression.

/*
switch expression {
case x:
  
case y:
   
case z:
...
default:
 
}
*/


// multi case
/*
	switch expression {
case x,y:
   
case v,w:
   
case z:
...
default:
 
}
*/

// switch with no condition

/*
t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
*/