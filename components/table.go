package components

import (
	"dinning-hall/util"
	"time"
)

type Table struct {
	TableId int
	State   State
	Order   Order
}

type State int

const (
	free State = iota
	waitingMakeOrder
	waitingServeOrder
)

func InitTable(tableId int) *Table {
	return &Table{
		TableId: tableId,
		State:   free,
	}
}

func WaitingMakeOrder(table *Table) *Table {
	defer time.Sleep(time.Duration(util.RandomizeNr(5)) * time.Second)
	return &Table{
		TableId: table.TableId,
		State:   waitingMakeOrder,
	}
}
func MakeOrder(table *Table) *Table {
	return &Table{
		TableId: table.TableId,
		State:   waitingServeOrder,
		Order:   CreateOrder(),
	}
}

func OrderServed(table *Table) *Table {
	return &Table{
		TableId: table.TableId,
		State:   free,
	}
}