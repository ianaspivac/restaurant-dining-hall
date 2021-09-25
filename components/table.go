package components
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
func GetState() {

}

func WaitingMakeOrder(table *Table) *Table {
	return &Table{
		TableId: table.TableId,
		State:   waitingMakeOrder,
	}
}
func MakeOrder(table *Table) *Table {
	return &Table{
		TableId: table.TableId,
		State:   waitingServeOrder,
		Order:   CreateOrder(table.TableId),
	}
}

func OrderServed(table *Table) *Table {
	return &Table{}
}
