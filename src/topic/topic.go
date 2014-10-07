package topic

import (
	"actor"
	"message"
)

type Topic struct {
	name string
	members []*actor.Actor
}

func New(name string) *Topic {
	var t = new(Topic)
	t.name = name
	return t
}

func (t *Topic) Send(m *message.Message) {
	if m.Recipient() != t.name {
		return
	}
	for i := 0; i < len(t.members); i++ {
		t.members[i].Send(m)
	}
}

func (t *Topic) Subscribe(a *actor.Actor) {
	t.members = append(t.members,a)	
}

func (t *Topic) Unsubscribe(a *actor.Actor) {
	for i,v := range t.members {
		if v == a {
			t.members = append(t.members[:i],t.members[i+1:]...)
			return
		}
	}
}

