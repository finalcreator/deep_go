package confinement

import (
	"fmt"
	"time"
)

func OwnershipA() {
	chanOwner := func() <-chan int {
		results := make(chan int, 5) // <1>
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
				time.Sleep(time.Millisecond*500)
			}
		}()
		return results
	}

	consumer := func(results <-chan int) { // <3>
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner() // <2>
	consumer(results)
}


func OwnershipB() {
	chanOwner := func() <-chan int {
		results := make(chan int, 5) // <1>
		go func() {
			//defer close(results)
			for i := 0; i <= 5; i++ {

				results <- i
				time.Sleep(time.Millisecond * 500)
			}
			 close(results)
			for i := 0; i <= 5; i++ {

				//results <- i
				fmt.Println(i)
				time.Sleep(time.Millisecond * 500)
			}
		}()
		return results
	}

	consumer := func(results <-chan int) { // <3>
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner() // <2>
	consumer(results)
}