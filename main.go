package main

import (
	component "dinning-hall/components"
	"dinning-hall/util"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	const nrTables int = 5
	var randMaximum int = 10000
	var table *component.Table
	table = component.InitTable(util.RandomizeNr(randMaximum))
	fmt.Printf("%+v\n", table)
	table = component.WaitingMakeOrder(table)
	fmt.Printf("%+v\n", table)
	table = component.MakeOrder(table)
	fmt.Printf("%+v\n", table)
}
