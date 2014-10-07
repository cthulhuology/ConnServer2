package actor

import (
	"message"
)

type handler func (a *Actor, m *message.Message) 

type Actor struct {
	name string
	inbox chan *message.Message
	outbox chan *message.Message
	does handler 
}

func New ( name string, does handler ) *Actor {
	var a = &Actor{}
	a.name = name
	a.does = does
	a.inbox = make(chan *message.Message)
	return a
}

func (a *Actor) Send (message *message.Message) {
	a.inbox <- message
}

func (a *Actor) Wait (outbox chan *message.Message) {
	a.outbox = outbox 
	for {
		m := <-a.inbox
		go a.does(a,m)
	}
}

