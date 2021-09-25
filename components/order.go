package components

type Order struct {
	OrderId int
	Priority int
	MenuItemIds []int
	MaxPreparationTime float32
}