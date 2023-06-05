package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	var intgroup []int
	fmt.Println("Hi")
	wg.Add(3)
	go func(*sync.WaitGroup, *sync.Mutex) {
		fmt.Println("One R")
		mut.Lock()
		intgroup = append(intgroup, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(*sync.WaitGroup, *sync.Mutex) {
		fmt.Println("Two R")
		mut.Lock()
		intgroup = append(intgroup, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(*sync.WaitGroup, *sync.Mutex) {
		fmt.Println("Three R")
		mut.Lock()
		intgroup = append(intgroup, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	wg.Wait()
	fmt.Println(intgroup)
}
