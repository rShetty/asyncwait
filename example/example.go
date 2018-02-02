package example

import "fmt"

func doWork(inChan chan int, outChan chan int) {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Hello", i)
			inChan <- i
			if i == 3 {
				outChan <- i
			}
		}
	}()
}
