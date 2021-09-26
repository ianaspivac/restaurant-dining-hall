package main

import (
	component "dinning-hall/components"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)
const nrTables int = 5
const nrWaiters int = 2
var tables [nrTables]*component.Table
var waiters [nrWaiters]*component.Waiter

func recieveOrder(c *gin.Context) {
	var order *component.OrderToSend
	if err := c.BindJSON(&order); err != nil {
		return
	}

	component.OrderServed(tables[order.TableId-1])
	fmt.Printf("Recieved prepared order: %+v \n",order)
	c.IndentedJSON(http.StatusCreated, order)
}

func main() {
	router := gin.Default()
	router.POST("/distribution", recieveOrder)

	rand.Seed(time.Now().UnixNano())


	for i := 0; i < nrTables; i++ {
		tables[i] = component.InitTable(i + 1)
	}
	for i := 0; i < nrWaiters; i++ {
		waiters[i] = component.InitWaiter()
	}


	go func() {
		component.WaitersSupervise(nrWaiters,nrTables,tables,waiters)
	}()
	router.Run("localhost:8080")
}
