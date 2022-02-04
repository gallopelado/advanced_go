package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Memory struct {
	f     Function
	cache map[int]FunctionResult
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

func (m *Memory) Get(key int /*, wg *sync.WaitGroup , lock *sync.RWMutex*/) (interface{}, error) {
	result, exists := m.cache[key]
	if !exists {
		//defer wg.Done()
		//lock.Lock()
		result.value, result.err = m.f(key)
		m.cache[key] = result
		//lock.Unlock()
	}
	return result.value, result.err
}

func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func procesarNumeros(n int, wg *sync.WaitGroup) {
	//var lock sync.RWMutex
	defer wg.Done()
	cache := NewCache(GetFibonacci)

	start := time.Now()
	value, err := cache.Get(n /*, wg , &lock*/)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%d, %s, %d\n", n, time.Since(start), value)
	//wg.Wait()
	//wg.Done()
}

func main() {
	var wg sync.WaitGroup
	fibo := []int{42, 40, 41, 42, 38}
	cantidadN := len(fibo)
	wg.Add(cantidadN)
	for _, n := range fibo {
		go procesarNumeros(n, &wg)
	}
	wg.Wait()
}
