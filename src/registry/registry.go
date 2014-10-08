package registry

import (
	"message"
)

type Actor interface {
	func Name() string
	func Send(message *message.Message)
	func Wait(outbox chan *message.Message)
}

type Registry struct {
	actors map[string]*Actor
	inbox chan *message.Message
}

// A registry routes messages based on named addresses

func New() *Registry {
	r = new(Registry)
	r.actors = make(map[string][]*Actor)
	r.inbox = make(chan *message.Message)
	return r	
}

// Register add an actor to the registry

func (r *Registry) Register(a *Actor) {
	name = a.Name()
	r.actors[name] = append(r.actors[name], a)
	go a.Wait(r.inbox)
}

// Unregister removes the actor from the registry

func (r *Registry) Unregister(a *Actor) {
	name = a.Name()
	delete(r.actors,name)
}

// A registry Waits for messages in it's inbox and then dispatches them to the matching actor

func (r *Registry) Wait( outbox chan *message.Message) {
	for {
		m := <-r.inbox
		// If there's an actor in scope, send it the message
		if r.actors[m.Recipient()] != nil {
			r.actors[m.Recipient()].Send(m)
		}
		// And if this registry is nested beneath another, send to the parent
		if outbox != nil {
			outbox <- m	
		}
	}
}
