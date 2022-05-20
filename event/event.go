package event

import (
	"dependency_injection_tut/greeter"
	"fmt"
)

type Event struct{
	Greeter greeter.Greeter
}

func NewEvent(g greeter.Greeter) Event{
	return Event{Greeter: g}
}


func (e Event) Start(){
	msg:=e.Greeter.Greet()
	fmt.Println(msg)
}
