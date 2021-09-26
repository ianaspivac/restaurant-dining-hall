package main

import (
	component "dinning-hall/components"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func recieveOrder(c *gin.Context) {
	var order *component.OrderToSend
	if err := c.BindJSON(&order); err != nil {
		return
	}

	component.GetTable(order.TableId-1).SetState(component.Free)
	fmt.Printf("Recieved prepared order: %+v \n",order)
	c.IndentedJSON(http.StatusCreated, order)
}

func main() {
	router := gin.Default()
	router.POST("/distribution", recieveOrder)

	rand.Seed(time.Now().UnixNano())

	component.InitTable()
	component.InitWaiter()
	component.TableSupervise()
	component.WaitersSupervise()

	router.Run(":8080")
}
