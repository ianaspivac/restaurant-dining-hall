package components

import (
	"bytes"
	"dinning-hall/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Waiter struct {
	ServedTable bool
	OrderToSend *OrderToSend
}

type OrderToSend struct {
	TableId            int     `json:"table_id"`
	OrderId            int     `json:"order_id"`
	Priority           int     `json:"priority"`
	MenuItemIds        []int   `json:"items"`
	MaxPreparationTime float32 `json:"max_wait"`
	PickUpTime         int64   `json:"pick_up_time"`
}
func InitWaiter()*Waiter{
	return &Waiter{
		ServedTable: false,
	}
}
func GetOrder(table *Table) *OrderToSend {
	return &OrderToSend{
		TableId:            table.TableId,
		OrderId:            table.Order.OrderId,
		Priority:           table.Order.Priority,
		MenuItemIds:        table.Order.MenuItemIds,
		MaxPreparationTime: table.Order.MaxPreparationTime,
		PickUpTime:         time.Now().Unix(),
	}
}
func UpdateWaiter(order *OrderToSend) *Waiter {
	return &Waiter{
		ServedTable: true,
		OrderToSend: order,
	}
}
func takingOrder(table *Table) (*Table, *Waiter) {
	var newTable = MakeOrder(table)
	var updateWaiter = UpdateWaiter(GetOrder(newTable))
	return newTable, updateWaiter
}
func WaitersSupervise(nrWaiters int, nrTables int, tables [5]*Table, waiters [2]*Waiter) {
	var getRandomTable int
	var wg sync.WaitGroup
	wg.Add(nrWaiters)
	var m sync.Mutex
	for i := 0; i < nrWaiters; i++ {
		go func(i int) {
			defer wg.Done()
			for {
				getRandomTable = util.RandomizeNr(nrTables - 1)
				if tables[getRandomTable].State == 0 {
					tables[getRandomTable] = WaitingMakeOrder(tables[getRandomTable])
				}
				m.Lock()
				for j := 0; j < nrTables; j++ {
					if tables[j].State == 1 && waiters[i].ServedTable == false {
						time.Sleep(1 * time.Second)
						tables[j], waiters[i] = takingOrder(tables[j])
						fmt.Printf("%+v\n", i)
						fmt.Printf("%+v\n", waiters[i].OrderToSend)

						jsonBody, err := json.Marshal(waiters[i].OrderToSend)
						if err != nil{
							log.Panic(err)
						}
						contentType:= "application/json"
						_, err = http.Post("http://localhost:8081/order", contentType, bytes.NewReader(jsonBody))
						if err != nil {
							return
						}
					}
				}
				m.Unlock()
			}
		}(i)
	}
	wg.Wait()
}

