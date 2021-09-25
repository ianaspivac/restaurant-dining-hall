package components

type Table struct {
	TableId int
	State State
	Order Order
}

type State int

const (
	free State = iota
	waitingMakeOrder
	waitingServeOrder
)

func GetState(){

}

func MakeOrder (){

}
func WaitingServeOrder(){

}
