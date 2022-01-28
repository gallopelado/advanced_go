package main

import (
	"fmt"
	"sync"
)

var balance = 100
var wg sync.WaitGroup
var lock sync.RWMutex

func main() {
	wg.Add(2)
	fmt.Println("Balance inicial=", balance)
	go depositar(100, &wg, &lock)
	go depositar(100, &wg, &lock)
	wg.Wait()
	fmt.Println("Balance final=", getBalance(&lock))
}

func depositar(monto int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	lock.Lock()
	balance = monto + balance
	lock.Unlock()
}

func getBalance(lock *sync.RWMutex) int {
	lock.RLock()
	b := balance
	lock.RUnlock()
	return b
}
