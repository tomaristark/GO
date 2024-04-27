package main

//GO Routines
//Concurrency in Go refers to the ability of a 
//Go program to execute multiple tasks 
//or operations concurrently, allowing
// different parts of the program to run independently 
//and potentially in parallel. Go provides built-in 
//features to support concurrency effectively, 
//making it easier to write concurrent programs.

//?Goroutines
//Goroutines are lightweight threads of execution managed by the Go runtime. 
//They enable concurrent execution of functions or methods. 
//You can start a new goroutine using the go keyword followed by a function call.

/*
go func() {

}()

*/

//Channels
// Channels are used to communicate and synchronize data between goroutines.
//They provide a way for goroutines to send and receive values. 
//Channels can be created using the make function.
/*
ch := make(chan int)
*/


//Channel Operations
//Channels support two main operations: sending (<-) and receiving (<-).
// Data is sent to a channel using the send operation, 
//and data is received from a channel using the receive operation.
/*
ch <- value // Send operation
result := <-ch // Receive operation
*/

//Buffered Channels
//Channels can be buffered, meaning they can hold a fixed number of elements.
// Sending to a buffered channel blocks only when the buffer is full, 
//and receiving blocks only when the buffer is empty.

/*ch := make(chan int, 10) // Buffered channel with capacity 10*/

//Select Statement
//The select statement allows you to wait on multiple channel operations simultaneously.
// It provides a way to handle multiple channels in a non-blocking manner
/*
select {
case <-ch1:
   / Handle operation from ch1
case <-ch2:
    / Handle operation from ch2
}
*/


//Mutexes and Locks:
//Go also provides synchronization primitives like mutexes (sync.Mutex) 
//to protect shared resources and prevent data races when accessing
// them from multiple goroutines

/*
var mutex sync.Mutex
mutex.Lock()

/Critical section
mutex.Unlock()

*/

//Goroutine
//In Go, a goroutine is a lightweight thread of execution.
// Goroutines enable concurrent programming, allowing functions to run 
//concurrently with other functions. They're very efficient in terms 
//of memory usage and overhead compared to traditional threads.

/*
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
/when you  run the program you will see world and hello will show 
randomly cause work concurrently
*/


//Channel
//Channels are a fundamental feature in Go for communication 
//and synchronization between goroutines. 
//They provide a way for goroutines to communicate with each other 
//and synchronize their execution. Channels allow you to pass data 
//between goroutines in a safe and efficient way, preventing race conditions and ensuring that data is properly synchronized.

//Creating channel

/*
ch := make(chan int) Creates a channel that can pass integers
*/
//Sending and Receiving data

/*
ch <- 42 // Sending data into the channel
value := <-ch // Receiving data from the channel
*/
//Channel Direction

//You can specify the direction of data flow in a channel. 
//By default, channels are bidirectional, 
//meaning they can be used for both sending and receiving.
// However, you can also create channels that are restricted to only sending 
//or only receiving data by specifying the direction when declaring the channel.

/* 
sendOnlyCh := make(chan<- int) // Channel that can only send integers
receiveOnlyCh := make(<-chan int) // Channel that can only receive integers
*/

//Blocking Operations

// Sending to or receiving from a channel will block the execution 
//of the sending or receiving goroutine until the other side is ready.
// This makes channels useful for synchronization between goroutines.

//Closing Channel

// You can close a channel using the close function. 
//This indicates that no more values will be sent on the channel.
// Receiving from a closed channel will return the zero value
// for the channel's type immediately

//You can check if a channel is closed
// using a comma-ok idiom when receiving from a channel

/*
close(ch)
value, ok := <-ch



*/
//value  will be 0 and ok is false when channel is closed

//!Note: Only the sender should close a channel, never the receiver.
//! Sending on a closed channel will cause a panic.
//!Channels aren't like files; 
//!you don't usually need to close them.
//! Closing is only necessary when the receiver must be told 
//!there are no more values coming, such as to terminate a range loop.

//Buffered Channels

//Buffered channels in Go allow you to store a limited number of values 
//without needing a corresponding receiver ready to read them immediately.
// This buffer acts as a queue, allowing senders to put values into the channel 
//without blocking until there's no more space in the buffer.

/*
Imagine you're at a small cafe with a limited number of tables.
When customers come in, they can sit at a table immediately if 
there are empty tables available. 
If all the tables are occupied, new customers have to wait 
until a table becomes available.
Now, think of the tables as slots in the buffer of a buffered channel,
and customers as values being sent into the channel. If there's space 
available in the buffer (empty tables), values can be sent into the channel 
without blocking. But if the buffer is full (all tables are occupied),
sending a new value into the channel will block until there's space available 
in the buffer (a table becomes empty).
*/

/*
package main

import "fmt"

func main() {
    /Create a buffered channel with a capacity of 2
    ch := make(chan int, 2)

    /Send values into the buffered channel
    ch <- 1
    ch <- 2

    /Attempting to send another value will block
    /until there's space available in the buffer
    ch <- 3 // This line will cause a deadlock if uncommented

    /Receive values from the buffered channel
    fmt.Println(<-ch) // Prints: 1
    fmt.Println(<-ch) // Prints: 2
}
*/

//Select
//The select statement lets a goroutine 
//wait on multiple communication operations.

//A select blocks until one of its cases can run, 
//then it executes that case. It chooses one at random 
//if multiple are ready.

//The select statement in Go is like a switch statement,
// but it's specifically designed for channels.
// It lets you wait on multiple channel operations simultaneously
//and performs whichever one is ready.

//Imagine you're waiting for a bus,
// a taxi, or a friend to pick you up. 
//You don't know which will arrive first, 
//but you're ready to go with whoever comes first. 
//The select statement works similarly

/*
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(2 * time.Second)
        ch1 <- "Hello"
    }()

    go func() {
        time.Sleep(1 * time.Second)
        ch2 <- "World"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Received:", msg1)
        case msg2 := <-ch2:
            fmt.Println("Received:", msg2)
        }
    }
}

*/

//In this example, there are two goroutines sending messages on two different channels (ch1 and ch2).
// The select statement waits for either channel to have data ready to receive.
// Whichever message arrives first will be printed, and the loop will continue until both channels have been read from.

//The default case in a select is run if no other case is ready.
//Use a default case to try a send or receive without blocking
/*
select {
case i := <-c:
    /use i
default:
    /receiving from c would block
}
*/

//Mutex and lock
//Mutually Exclusive
//In concurrent programming, multiple goroutines accessing 
//shared resources can lead to race conditions, 
//where the outcome of the program depends on the timing 
//of their execution. To prevent this, you can use synchronization
// primitives like sync.Mutex, which provides mutual exclusion, 
//ensuring that only one goroutine can access a shared resource at a time


// /Imagine you have a variable representing the balance of a bank account, 
//and you want to ensure that withdrawals and deposits to this account 
//are executed safely, without risking inconsistencies due to concurrent access.

/*
package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance int
	mutex   sync.Mutex
}

func (acc *BankAccount) Deposit(amount int) {
	acc.mutex.Lock() //when use bankAccount lock the value
	defer acc.mutex.Unlock() //when finish unlock the value
	acc.balance += amount
}

func (acc *BankAccount) Withdraw(amount int) {
	acc.mutex.Lock()
	defer acc.mutex.Unlock()
	acc.balance -= amount
}

func (acc *BankAccount) Balance() int {
	acc.mutex.Lock()
	defer acc.mutex.Unlock()
	return acc.balance
}

func main() {
	account := BankAccount{}

	account.Deposit(100)
	fmt.Println("Balance after deposit:", account.Balance()) // Should print 100

	account.Withdraw(50)
	fmt.Println("Balance after withdrawal:", account.Balance()) // Should print 50
}

*/

//Wait groups, represented by sync.WaitGroup in Go,
// provide a simple and effective way to synchronize 
//the execution of multiple goroutines.
// They allow you to wait for a collection of goroutines 
//to finish executing before proceeding with further actions.
//WaitGroup is like a concurrent-safe count tracker.
// When you call Add it increments to count and when you call Done is decrements it.
// Then you place a call to Wait which blocks 
//until the counter is zero and the main function can exist. 
/*
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Decrement the wait group counter when the goroutine finishes
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup // Create a wait group

    for i := 1; i <= 3; i++ {
        wg.Add(1) // Increment the wait group counter for each goroutine
        go worker(i, &wg)
    }

    wg.Wait() // Wait for all goroutines to finish
    fmt.Println("All workers have completed their tasks")
}

*/

//-We create a wait group wg using var wg sync.WaitGroup.
//-Inside the loop, we launch three goroutines to simulate 
//workers performing tasks concurrently.
// Before launching each goroutine, 
//we increment the wait group counter using wg.Add(1).
//-Each goroutine runs the worker function, 
//which simulates performing some work and then calls wg.Done()
// to decrement the wait group counter when finished.
//-Finally, we call wg.Wait() to block the main goroutine 
//until all workers have completed their tasks.
//Wait groups provide a convenient mechanism 
//for coordinating the execution of multiple 
//goroutines and waiting for them to finish before 
//continuing with further actions. They are commonly used 
// in Go for managing concurrent tasks and ensuring synchronization.