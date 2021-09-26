package components

import (
	"dinning-hall/util"
)

type Order struct {
	OrderId            int
	Priority           int
	MenuItemIds        []int
	MaxPreparationTime float32
}

func CreateOrder() Order {

	menuItemIds, maxPrepTime := getMenuItems()
	return Order{
		OrderId:            util.RandomizeNr(10000),
		MenuItemIds:        menuItemIds,
		MaxPreparationTime: maxPrepTime,
		Priority: 1,
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
