/* topic

	The topic package provides a way to match messages with individual actors using the message
	topic metadata.  This is handy when a number of actors need to respond to a subset of messages
	coordinated by some common concept.  For example, one might produce a topic string:

		"mail usage ord"

	And another server may publish to a related topic:
	
		"mail usage dfw"

	Where the two topics are related by both their primary topic (mail), their secondary topic (usage),
	and differ in their tertiary topic (location: chicago or dallas).  One could easily see three
	different types of monitoring actors that might be interested in this sort of data:

		* mail monitoring - looks at all mail messages across datacenters
		* usage monitoring - looks at all application usage across datacenters
		* datacenter monitoring - looks at behavior within a datacenter
	
	Each of these different types of actors could use topic routers to listen for applicable information:

		* mail monitoring - "mail .*"
		* usage monitoring - "\w+ usage .*"
		* datacenter monitoring (chicago only) - "\w+ \w+ ord"	
	
	The topic router can provide for a wide variety of applications, and a single topic router could
	accomodate all of these use cases without needing additional logic.
*/

package topic

import (
	"regexp"
	"actor"
	"message"
)

// The Topic type provides a name (for the Registry), and two parallel arrays of members and bindings.
// In general I dislike parallel arrays, but in this case I'm doing it to avoid an additional type.
// Future structural changes may occur to clean it up if this proves painful.

type Topic struct {
	name string	
	members []*actor.Actor
	bindings []*regexp.Regexp
}


// New is the default constructor for topic routers, and merely requires a name so that it can 
// be found in the registry

func New(name string) *Topic {
	var t = new(Topic)
	t.name = name
	return t
}


// Name allows a Topic to reflect upon it's registered name.  Other actors may send messages to it
// by Name() as long as they are routed through the registry.

func (t * Topic) Name () string {
	return t.name
}


// Send does what is says to a message to each of the topic router's members whose bindings match the
// message's Topic() metadata field.  If the match fails, then no message is sent to that member actor.

func (t *Topic) Send(m *message.Message) {
	for i := 0; i < len(t.members); i++ {
		if t.bindings[i].MatchString(m.Topic()) {
			t.members[i].Send(m)
		}
	}
}


// Subscribe adds a new member to the topic router, registering the associated regexp with the actor.
// It is possible to register a single actor to the topic router multiple times with different regexps
// which allows an actor to receive multiple subtopics from a given message stream.

func (t *Topic) Subscribe(a *actor.Actor, r *regexp.Regexp) {
	t.members = append(t.members,a)	
	t.bindings = append(t.bindings,r)
}

// Unsubscribe removes all instances of an actor from a topic router.  All of the associated bindings
// are removed as well.  If useful, this method may take an additional regexp in the future so that only
// specific bindings may be removed.  This may require additional attention to the datastructure.

func (t *Topic) Unsubscribe(a *actor.Actor) {
	for i,v := range t.members {
		if v == a {
			t.members = append(t.members[:i],t.members[i+1:]...)
			t.bindings = append(t.bindings[:i],t.bindings[i+1:]...)
		}
	}
}

