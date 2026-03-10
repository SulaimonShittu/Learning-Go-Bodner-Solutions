package main

import (
	"context"
	"fmt"
	"math"
	"sync"
)

func main() {
	// E-01
	goLaunch()
	// E-02
	goLaunch2()
	// E-03
	funcCache := sync.OnceValue(func() map[int]float64 {
		data := mapBuild()
		return data
	})
	data := funcCache()
	x := 1000
	y := 1
	for x < len(data) {
		fmt.Printf("The %v 1000th value is : %v\n", y, data[x])
		x += 1000
		y += 1
	}
}

func goLaunch() {
	var done int
	msgChan := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	wgp := &sync.WaitGroup{}
	wgp.Add(1)
	go func() {
		for i := range 10 {
			select {
			case msgChan <- i + 1:
			case <-ctx.Done():
				return
			}
		}
		done += 1
		if done == 2 {
			cancel()
		}
	}()
	go func() {
		for i := range 10 {
			select {
			case msgChan <- (i + 1) * 2:
			case <-ctx.Done():
				return
			}
		}
		done += 1
		if done == 2 {
			cancel()
		}
	}()
	go func() {
		defer wgp.Done()
		for {
			select {
			case val := <-msgChan:
				fmt.Printf("The available value is : %v\n", val)
			case <-ctx.Done():
				return
			}
		}
	}()
	wgp.Wait()
}

func goLaunch2() {
	var done int
	msg, msg2 := make(chan int), make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		for i := range 10 {
			select {
			case msg <- i:
			case <-ctx.Done():
				return
			}
		}
		done += 1
		if done == 2 {
			cancel()
		}
	}()
	go func() {
		for i := range 10 {
			select {
			case msg2 <- i * 2:
			case <-ctx.Done():
				return
			}
		}
		done += 1
		if done == 2 {
			cancel()
		}
	}()

	for {
		select {
		case val := <-msg:
			fmt.Printf("The value: %v is from the first gorountine\n", val)
		case val := <-msg2:
			fmt.Printf("The value: %v is from the second gorountine\n", val)
		case <-ctx.Done():
			return
		}
	}
}

func mapBuild() map[int]float64 {
	data := map[int]float64{}
	for i := range 100_000 {
		data[i] = math.Sqrt(float64(i))
	}
	return data
}
