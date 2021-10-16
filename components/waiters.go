package components

import (
	"bytes"
	"dinning-hall/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const NrWaiters int = 2

var Waiters [NrWaiters]*Waiter

type Waiter struct {
	id           int
	waiterTables []*Table
}

type OrderToSend struct {
	Order

	TableId    int   `json:"table_id"`
	WaiterId   int   `json:"waiter_id"`
	PickUpTime int64 `json:"pick_up_time"`
}

func InitWaiter() {
	for idx, _ := range Waiters {
		Waiters[idx] = &Waiter{
			id:           idx+1,
			waiterTables: tables,
		}
	}
}

func (w *Waiter) Supervise() {
	for {
		for _, val := range w.waiterTables {
			if val.GetState() == WaitingMakeOrder {

				val.MakeOrder()
				time.Sleep(time.Duration(util.RandomizeNr(4)) * time.Second)

				order := val.GetOrder()

				fmt.Printf("Waiter %d picked order from table %d: %+v\n", w.id, val.TableId,order)

				sendOrder := &OrderToSend{
					Order: order,

					TableId:    val.TableId,
					WaiterId:   w.id,
					PickUpTime: time.Now().Unix(),
				}

				jsonBody, err := json.Marshal(sendOrder)
				if err != nil {
					log.Panic(err)
				}

				contentType := "application/json"
				_, err = http.Post("http://localhost:8081/order", contentType, bytes.NewReader(jsonBody))
				if err != nil {
					log.Panic(err)
				}
			}
		}
	}
}

func WaitersSupervise() {
	for idx, _ := range Waiters {
		go Waiters[idx].Supervise()
	}
}
