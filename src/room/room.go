/* room

	The room package provides a basic chat room style interface for fanout message routing.
	Each member of the room receives all messages sent to the room, and can see all of the
	members of that room.  On changes in room state, the room also announces major events
	like the addition or removal of new members.

	Primarily, rooms are use to share message flows between logically proximal actors.  Rather
	than filtering the message flows, the room acts as a hub and simply provides a named rebroadcast
	point for a group of actors.  It is also possible to use a room to coordinate a collection of
	processes, such that each of the members of the room can coordinate a share of the work.

	For example, a room router may be subscribed to a topic router to provide a group of workers
	with commands that match the given topic.  These workers all know about the existence of the
	other workers, and can safely divy up the work with a manager keeping track of how many workers
	should be in the room for a given point in time workload.
	
	Another simple use for a room is zone chat and interactions within a game, where the proximity
	of any two users is dependant upon which room they occupy.  Since rooms can contain other rooms,
	it is possible to create hiearchies of rooms to break down regions into smaller and smaller
	partitions.

*/
package room

import (
	"actor"
	"message"
)

// The Room struct keeps track of both a registered name of the room (as seen in the registry) and a
// list of actors who belong in the room.  

type Room struct {
	name string
	members []*actor.Actor	
}

// The New constructor simply takes a name for the registration of the room, and starts with no members

func New(name string) *Room {
	var r = new(Room)
	r.name = name
	return r
}

// Send does what is says to a given message to each of the members of the room.

func (r *Room) Send(m *message.Message) {
	// A room sends too all of the members in the room, they can choose to listen on their end
	for i := 0; i < len(r.members); i++ {
		r.members[i].Send(m)
	}
}

// Join places an actor in the room adding it to its members, and also announces the event to the inhabitants
// Any given actor may only join a room once, and an actor may not join a room multiple times.

func (r *Room) Join(a *actor.Actor) {
	// Prevent a member of a room from joining again, by iterating over the members
	for i,v := range r.members {
		// and if we find the member already there
		if v == a {
			// we exit immediately
			return
		}
	}
	// otherwise we append the new member to the membership
	r.members = append(r.members,a)
	// and announce that they have joined the room
	r.Send(message.New(`["announce","join",a.Name()]`))
}

// Leave removes an actor from the room and announces that the actor has left the stage.
// Should an actor some how manage to get on stage multiple times, all occurances will be removed
// NB: this should only be possible if someone violated the Join logic above!!!

func (r *Room) Leave(a *actor.Actor) {
	// When we get a request to leave a room, we iterate over all of the members
	for i,v := range r.members {
		// and if we find the member we're looking for
		if v == a {
			// we splice them out of the member list
			r.members = append(r.members[:i],r.members[i+1:]...)
			// and annouce that they have left the room to the others inside
			r.Send(message.New(`["announce","leave",a.Name()]`))	
		}
	}
}

// Members reports a list of actor names that currently occupy the room.  Should an actor not have a 
// registered name, then that actor will be ommitted from the membership list.  Nameless actors are 
// reserved for bots and other magical creatures which exist only in the system and not in the public model.

func (r *Room) Members() []string {
	members []string
	for i,v := range r.members {
		members = append(members,v.Name())
	}
	return members
}


