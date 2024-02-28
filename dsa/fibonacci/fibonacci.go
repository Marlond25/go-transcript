package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fibonacci(n int) (int, error) {

	time.Sleep(time.Second)

	if n < 0 {
		return n, errors.New("'n' must be a positive integer! \n")
	}

	if n == 0 {
		return 0, nil
	}

	if n == 1 {
		return 1, nil
	}

	fib := 1
	fmin2 := 0
	fmin1 := 1

	for i := 2; i <= n; i++ {
		fib = fmin1 + fmin2
		fmin2 = fmin1
		fmin1 = fib
	}

	return fib, nil

}

func fibonacci2(n int) {

	time.Sleep(time.Second)

	defer wg.Done()

	if n < 0 {
		fmt.Printf("'n' must be a positive integer 2! \n")
		return
	}

	if n == 0 {
		fmt.Printf("fibonacci2(%d) = %v \n", n, 0)
		return
	}

	if n == 1 {
		fmt.Printf("fibonacci2(%d) = %v \n", n, 1)
		return
	}

	fib := 1
	fmin2 := 0
	fmin1 := 1

	for i := 2; i <= n; i++ {
		fib = fmin1 + fmin2
		fmin2 = fmin1
		fmin1 = fib
	}

	fmt.Printf("fibonacci2(%d) = %v \n", n, fib)

}

func main() {

	var start = time.Now()

	count := 100

	for i := -1; i < count; i++ {

		res, err := fibonacci(i)

		if err != nil {
			fmt.Printf("%s", err)
		} else {
			fmt.Printf("fibonacci(%d) = %v \n", i, res)
		}

	}

	var duration = time.Now().Sub(start).Seconds()

	time.Sleep(time.Second * 3)

	var start2 = time.Now()

	for j := 1; j < count+2; j++ {

		param := j - 2

		wg.Add(1)

		go fibonacci2(param)
	}

	wg.Wait()

	var duration2 = time.Now().Sub(start2).Seconds()

	fmt.Println(duration)
	fmt.Println(duration2)

}
