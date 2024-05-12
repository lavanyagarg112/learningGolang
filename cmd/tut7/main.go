package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

var results =  []string{}  // is this a slice??

func main(){
	fmt.Println("Goroutines")

	// way to launch multiple functions
	// and launch them concurrently
	// not same as parallelism
	// just run back and forth between the tasks
	// while waiting for one task, work on diff tasks

	// multiple cores achieves parallelism... but still concurrent, each core

	// multicore cpu -> parallelism

	t0 := time.Now()
	for i := 0; i <len(dbData); i++{
		// dbCall(i) // run simultaneously
		// go dbCall(i) // executed them in the background but exited before executing them

		// solution: use weight groups!

		wg.Add(1)
		go dbCall(i)

		// we see diff order of execution! and much faster!
	}
	wg.Wait() // wait for the counter to reach 0
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))


}

func dbCall(i int) {
	// simulate db call delay
	var delay float32 = rand.Float32() * 2000 
	// if we remove this random delay just do 2000
	// then if we append our results into another array and print it in main
	// we see that some of the stuff will work together on the same memory location
	// leading to missing data!
	// our array might only have id3, id4, id1!
	// solution: mutex! mututal execution

	// m.Lock() // putting it here defeats the purpose of concurrency!
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	m.Lock() // the goroutine waits until the lock is released, and sets it again
	results = append(results, dbData[i])
	m.Unlock()
	wg.Done() // decrements the counter

	// one drawback of mutex is it entirely locks out other go routines

	// we also have RWMutex (read write mutex)
	// reading results we can add a read lock - m.RLock(), m.RUnlock() - we are not reading while results are being written to, i.e. while Lock()
	// but if just RLock then can acquire Rlock and proceed.
	// a go routine can have multiple readlocks
	// it can only block code execution onto the m.Lock()
	// when hits m.Lock, all other locks must be cleared just like before
	// this is to prevent any access to the slice while writing etc
	// writing we can add write lock - m.Lock()


	// without RWMutex, even read was happening one at a time
}


// but if we have a task which counts to a 100000, and go routines need to do some work, which takes some work
// concurrency not so useful


// with dbCall we get constant time
// count is linear time based on amount of cores in the cpu