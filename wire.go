//go:build wireinject
// +build wireinject

package main

import (
	"dependency_injection_tut/event"
	"dependency_injection_tut/greeter"
	"dependency_injection_tut/message"

	"github.com/google/wire"
)

func InitializeEvent() event.Event{
	wire.Build(message.NewMessage,greeter.NewGreeter,event.NewEvent)
	return event.Event{}
}