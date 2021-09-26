package main

import (
	component "dinning-hall/components"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)



func main() {
	router := gin.Default()

	rand.Seed(time.Now().UnixNano())
	const nrTables int = 5
	const nrWaiters int = 2
	var tables [nrTables]*component.Table
	var waiters [nrWaiters]*component.Waiter

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
