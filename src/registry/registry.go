package registry

import (
	"actor"
	"message"
	"topic"
	"room"
)

type Registry struct {
	actors []*actor.Actor
	topics []*topic.Topic
	rooms []*room.Room
}

func New () {
	

}
