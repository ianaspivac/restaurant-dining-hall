package components

import (
	"sync"
	"time"
)

const NrTables int = 5
const (
	Free State = iota
	WaitingMakeOrder
	WaitingServeOrder
)

var tablesMutex sync.Mutex
var tables []*Table

type State int

type Table struct {
	TableId int

	stateMutex sync.Mutex
	state      State

	orderMutex sync.Mutex
	order      Order
}

func (t *Table) GetState() State {
	t.stateMutex.Lock()
	defer t.stateMutex.Unlock()

	return t.state
}

func (t *Table) SetState(newState State) {
	t.stateMutex.Lock()
	defer t.stateMutex.Unlock()

	t.state = newState
}

func (t *Table) GetOrder() Order {
	t.orderMutex.Lock()
	defer t.orderMutex.Unlock()

	return t.order
}

func (t *Table) SetOrder(newOrder Order) {
	t.orderMutex.Lock()
	defer t.orderMutex.Unlock()

	t.order = newOrder
}

func (t *Table) MakeOrder() {
	t.SetState(WaitingServeOrder)
	t.SetOrder(CreateOrder())
}

func (t *Table) OrderServed() {
	t.SetState(Free)
}

func InitTable(){
	for idx := 0; idx < NrTables; idx++ {
		tables = append(tables, &Table{
			TableId: idx+1,
			state: Free,
		})
	}
}

func TableSupervise() {
	go generateOrders()
}

func generateOrders(){
	for {
		for idx, _ := range tables {
			val := GetTable(idx)
			if val.GetState() == Free {
				val.SetState(WaitingMakeOrder)
			}
		}

		time.Sleep(time.Second * 3)
	}
}

func GetTable(idx int) *Table {
	tablesMutex.Lock()
	defer tablesMutex.Unlock()
	return tables[idx]
}

