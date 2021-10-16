package components

import (
	"dinning-hall/util"
)

type Order struct {
	OrderId            int     `json:"order_id"`
	Priority           int     `json:"priority"`
	MenuItemIds        []int   `json:"items"`
	MaxPreparationTime float32 `json:"max_wait"`
}

type CookFood struct {
	FoodId int `json:"food_id"`
	CookId int `json:"cook_id"`
}

type OrderPrepared struct {
	Order
	WaiterId       int        `json:"waiter_id"`
	TableId        int        `json:"table_id"`
	PickUpTime     int64      `json:"pick_up_time"`
	CookingTime    int64      `json:"cooking_time"`
	CookingDetails []CookFood `json:"cooking_details"`
}

func CreateOrder() Order {
	menuItemIds, maxPrepTime := getMenuItems()
	return Order{
		OrderId:            util.RandomizeNr(10000),
		MenuItemIds:        menuItemIds,
		MaxPreparationTime: maxPrepTime,
		Priority:           util.RandomizeNr(5),
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
	var quantityFoods int = util.RandomizeNr(5)
	var foodsId []int
	var maxPrepTime float32

	for i := 0; i <= quantityFoods; i++ {
		foodsId = append(foodsId, util.RandomizeNr(10))
	}
	maxPrepTime = GetMaxPrepTime(foodsId, quantityFoods)

	return foodsId, maxPrepTime
}
