package main

//Go is a concurrent  language but
//Go  supports parallelism through the use of multiple CPU cores.
// Go's runtime scheduler can automatically distribute goroutines across CPU cores 
//to achieve parallel execution.

///to know the number of cores in cpu go has the function called NumCPU
//var numCPU = runtime.NumCPU return the number  of cores

//runtime.GOMAXPROCS()
//This function sets the maximum number of operating system threads that can execute Go code simultaneously
//By default, Go uses only one CPU core, but you can increase this value to utilize multiple CPU cores, potentially achieving parallelism
// Setting GOMAXPROCS affects the behavior of the Go scheduler and the degree of concurrency in your program.

/*
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	/Set GOMAXPROCS to utilize multiple CPU cores
	runtime.GOMAXPROCS(2) // Adjust this value based on your system configuration

	/Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	/Define the number of goroutines to run
	numGoroutines := 5

	/Add the number of goroutines to the WaitGroup
	wg.Add(numGoroutines)

	/Launch multiple goroutines
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Println("Goroutine", id, "started")
			/Simulate some computation
			for j := 0; j < 100000000; j++ {
				/Some computation
			}
			fmt.Println("Goroutine", id, "finished")
		}(i)
	}

	/Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All goroutines have finished")
}

*/

//Using numCPU with GOMAXPROCS

/*
package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    /Get the number of available CPU cores
    numCPU := runtime.NumCPU()
    fmt.Println("Number of CPU cores available:", numCPU)

   / Set GOMAXPROCS to utilize all available CPU cores
    maxProcsBefore := runtime.GOMAXPROCS(numCPU)
    fmt.Println("GOMAXPROCS before setting:", maxProcsBefore)

    /Simulate CPU-bound work with goroutines
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Printf("Goroutine %d is running\n", id)
            for j := 0; j < 100000000; j++ {
                /Simulate some computation
            }
        }(i)
    }
    wg.Wait()

    fmt.Println("All goroutines have finished.")

    /Reset GOMAXPROCS to its original value
    runtime.GOMAXPROCS(maxProcsBefore)
    fmt.Println("GOMAXPROCS after resetting:", runtime.GOMAXPROCS(0))
}

*/