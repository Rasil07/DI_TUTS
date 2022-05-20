package greeter

import (
	"dependency_injection_tut/message"
)

type Greeter struct{
	Message message.Message
}


func NewGreeter(m message.Message) Greeter{
return Greeter{
	Message: m,
}
}

func (g Greeter) Greet() message.Message{
	return g.Message
}