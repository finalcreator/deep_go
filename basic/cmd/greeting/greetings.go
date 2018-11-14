package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) >= 2 {
		repeatNum, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < repeatNum; i++ {
			fmt.Printf("The no. is %d\n", i+1)
		}
	}

}
