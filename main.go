package main

import (
	component "dinning-hall/components"
	"dinning-hall/util"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func takingOrder(table *component.Table) (*component.Table, *component.Waiter) {
	var newTable = component.MakeOrder(table)
	var updateWaiter = component.UpdateWaiter(component.GetOrder(newTable))
	return newTable, updateWaiter
}
func main() {
	rand.Seed(time.Now().UnixNano())
	const nrTables int = 5
	const nrWaiters int = 2
	var tables [nrTables]*component.Table
	var waiters [nrWaiters]*component.Waiter

	for i := 0; i < nrTables; i++ {
		tables[i] = component.InitTable(i + 1)
	}
	for i := 0; i < nrWaiters; i++ {
		waiters[i] = component.InitWaiter()
	}

	var getRandomTable int
	var wg sync.WaitGroup
	wg.Add(nrWaiters)
	for i := 0; i < nrWaiters; i++ {
		go func(i int) {
			defer wg.Done()
			for {
				getRandomTable = util.RandomizeNr(nrTables - 1)
				if tables[getRandomTable].State == 0 {
					tables[getRandomTable] = component.WaitingMakeOrder(tables[getRandomTable])
				}
				var m sync.Mutex
				for j := 0; j < nrTables; j++ {
					m.Lock()
					if tables[j].State == 1 && waiters[i].ServedTable == false {
						time.Sleep(1 * time.Second)
						tables[j], waiters[i] = takingOrder(tables[j])
						fmt.Printf("%+v\n", i)
						fmt.Printf("%+v\n", waiters[i].OrderToSend)
					}
					m.Unlock()
				}
			}
		}(i)
	}
	wg.Wait()

}
