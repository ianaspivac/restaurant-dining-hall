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
	WaiterId           int     `json:"waiter_id"`
	Priority           int     `json:"priority"`
	MenuItemIds        []int   `json:"items"`
	MaxPreparationTime float32 `json:"max_wait"`
	PickUpTime         int64   `json:"pick_up_time"`
}

func InitWaiter() *Waiter {
	return &Waiter{
		ServedTable: false,
	}
}

func GetOrder(table *Table, waiterId int) *OrderToSend {
	return &OrderToSend{
		TableId:            table.TableId,
		OrderId:            table.Order.OrderId,
		WaiterId:          waiterId,
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
func takingOrder(table *Table, waiterId int) (*Table, *Waiter) {
	var newTable = MakeOrder(table)
	var updateWaiter = UpdateWaiter(GetOrder(newTable, waiterId))
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
				m.Lock()
				getRandomTable = util.RandomizeNr(nrTables)
				if tables[getRandomTable-1].State == 0 {
					tables[getRandomTable-1] = WaitingMakeOrder(tables[getRandomTable-1])
					fmt.Printf("Table %v is waiting  \n",tables[getRandomTable-1].TableId)
				}
				m.Unlock()
				m.Lock()
				for j := 0; j < nrTables; j++ {
					if tables[j].State == 1 {
						time.Sleep(1 * time.Second)
						tables[j], waiters[i] = takingOrder(tables[j], i+1)
						fmt.Printf("Waiter %v picked order\n",i)
						fmt.Printf("%+v\n", waiters[i].OrderToSend)

						jsonBody, err := json.Marshal(waiters[i].OrderToSend)
						if err != nil {
							log.Panic(err)
						}
						contentType := "application/json"
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
