package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"deepgou.com/google_deep_go/example/loadBalancer/smartbalancer"
)

func main() {

	var insts []*smartbalancer.Instance
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := smartbalancer.NewInstance(host, 8080)
		insts = append(insts, one)
	}

	var balanceName = "random"
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}

	for {
		inst, err := smartbalancer.DoBalance(balanceName, insts)
		if err != nil {
			//fmt.Println("do balance err:", err)
			fmt.Fprintf(os.Stdout, "do balance err\n")
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}

}
