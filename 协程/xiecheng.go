package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan = make(chan int)

func Shopping(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s正在购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s购物结束\n", name)

	moneyChan <- money

	wait.Done()
}

func main() {
	var wait sync.WaitGroup
	nowtime := time.Now()
	wait.Add(3)
	go Shopping("牛哥", 3, &wait)
	go Shopping("牛爷爷", 4, &wait)
	go Shopping("牛小弟", 5, &wait)

	go func() {
		wait.Wait()
		close(moneyChan)
	}()

	// for {
	// 	money, ok := <-moneyChan
	// 	fmt.Println(money, ok)
	// 	if !ok {
	// 		break
	// 	}
	// }
	//两个for等价
	var moneyList []int

	for money := range moneyChan {
		fmt.Println(money)
		moneyList = append(moneyList, money)
	}

	fmt.Println("TIME:", time.Since(nowtime), "MONEYLIST:", moneyList)
}
