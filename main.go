package main

import (
	component "dinning-hall/components"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func takingOrder(table *component.Table) (*component.Table, *component.OrderToSend) {
	var newTable = component.MakeOrder(table)
	return newTable, component.GetOrder(newTable)

}
func main() {
	rand.Seed(time.Now().UnixNano())
	const nrTables int = 5
	var tables [nrTables]*component.Table

	for i := 0; i < nrTables; i++ {
		tables[i] = component.InitTable(i + 1)
		//fmt.Printf("%+v\n", tables[i])
	}
	var wg sync.WaitGroup
	wg.Add(nrTables)

	for i := 0; i < nrTables; i++ {
		go func(i int) {
			defer wg.Done()
			tables[i] = component.WaitingMakeOrder(tables[i])
			fmt.Printf("%+v\n", tables[i])
		}(i)
	}
	wg.Wait()

	//var orderToSend *component.OrderToSend

	//table, orderToSend = takingOrder(table)

	//fmt.Printf("%+v\n", orderToSend)

}
