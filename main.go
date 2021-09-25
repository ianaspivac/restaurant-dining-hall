package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Food struct {
	id               int
	name             string
	preparationTime  float32
	complexity       int
	cookingApparatus string
}

var Menu = []Food{
	{
		id:               1,
		name:             "pizza",
		preparationTime:  20,
		complexity:       2,
		cookingApparatus: "oven",
	},
	{
		id:               2,
		name:             "salad",
		preparationTime:  10,
		complexity:       1,
		cookingApparatus: "",
	},
	{
		id:               3,
		name:             "zeama",
		preparationTime:  7,
		complexity:       1,
		cookingApparatus: "stove",
	},
	{
		id:               4,
		name:             "Scallop Sashimi with Meyer Lemon Confit",
		preparationTime:  32,
		complexity:       3,
		cookingApparatus: "",
	},
	{
		id:               5,
		name:             "Island Duck with Mulberry Mustard",
		preparationTime:  35,
		complexity:       3,
		cookingApparatus: "oven",
	},
	{
		id:               6,
		name:             "Waffles",
		preparationTime:  10,
		complexity:       1,
		cookingApparatus: "stove",
	},
	{
		id:               7,
		name:             "Aubergine",
		preparationTime:  20,
		complexity:       2,
		cookingApparatus: "",
	},
	{
		id:               8,
		name:             "Lasagna",
		preparationTime:  30,
		complexity:       2,
		cookingApparatus: "oven",
	},
	{
		id:               9,
		name:             "Burger",
		preparationTime:  15,
		complexity:       1,
		cookingApparatus: "oven",
	},
	{
		id:               10,
		name:             "Gyros",
		preparationTime:  15,
		complexity:       1,
		cookingApparatus: "",
	},
}

type Order struct {
	TableId            int     `json:"table_id"`
	OrderId            int     `json:"order_id"`
	Priority           int     `json:"priority"`
	MenuItemIds        []int   `json:"items"`
	MaxPreparationTime float32 `json:"max_wait"`
	PickUpTime         int64   `json:"pick_up_time"`
}

func CreateOrder(tableId int) Order {

	menuItemIds, maxPrepTime := getMenuItems()
	return Order{
		TableId:            tableId,
		OrderId:            RandomizeNr(10000),
		MenuItemIds:        menuItemIds,
		MaxPreparationTime: maxPrepTime,
		PickUpTime:         time.Now().Unix(),
	}
}
func findMaxPrepTime(arr []float32) (max float32) {
	max = arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max * 1.3
}
func GetMaxPrepTime(foodsId []int, quantityFoods int) float32 {
	var maxPrepTimeFoods []float32
	for i := 0; i <= quantityFoods; i++ {
		maxPrepTimeFoods = append(maxPrepTimeFoods, Menu[foodsId[i]-1].preparationTime)
	}
	return findMaxPrepTime(maxPrepTimeFoods)
}

func getMenuItems() ([]int, float32) {
	var quantityFoods int = RandomizeNr(5)
	var foodsId []int
	var maxPrepTime float32

	for i := 0; i <= quantityFoods; i++ {
		foodsId = append(foodsId, RandomizeNr(10))
	}
	maxPrepTime = GetMaxPrepTime(foodsId, quantityFoods)

	return foodsId, maxPrepTime
}

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
func Randomizer() {
	var res bool
	min := 1
	max := 100
	if (rand.Intn(max-min+1)+min)%2 == 0 {
		res = false
	} else {
		res = true
	}
	fmt.Println(res)
}
func RandomizeNr(maximum int) int {
	min := 1
	max := maximum
	return rand.Intn(max-min+1) + min
}

func main() {
	rand.Seed(time.Now().UnixNano())
	const nrTables int = 5
	var randMaximum int = 10000
	var table *Table
	table = InitTable(RandomizeNr(randMaximum))
	fmt.Printf("%+v\n", table)
	table = WaitingMakeOrder(table)
	fmt.Printf("%+v\n", table)
	table = MakeOrder(table)
	fmt.Printf("%+v\n", table)
}
