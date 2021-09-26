package components

import "time"

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

