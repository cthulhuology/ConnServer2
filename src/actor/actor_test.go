package actor

import (
	"testing"
	"message"
)

func TestNew (t *testing.T) {
	var doit handler = func (a *Actor, m *message.Message) {
		t.Logf("Got message",m)
	}
	var act = New("dave",doit)
	if act.name != "dave" {
		t.Errorf("Invalid actor name %v", act.name)
	}	
	if act.does == nil  {
		t.Errorf("Invalid actor handler %v", act.does)
	}
}

func TestSend (t *testing.T) {
	var doit handler = func (a *Actor, m *message.Message) {
		t.Logf("Got message %v",m)
		if m.String() != `["hello","world"]` {
			t.Errorf("Failed to get hello world message! %v", m)	
		}
		a.outbox <- message.New(`["ok"]`)
	}
	var act = New("dave",doit)
	var result = make(chan *message.Message)
	go act.Wait(result)
	act.Send(message.New(`[ "hello", "world" ]`))
	t.Logf("responds %v", <- result)
}

func TestWait (t *testing.T) {
	var count = 0
	var doit handler = func (a *Actor, m *message.Message) {
		if m.String() != `["hello","world"]` {
			t.Errorf("Failed to get hello world message! %v", m)	
		}
		count += 1
		t.Logf("Received %v messages",count)
		if count == 3 {
			a.outbox <- message.New(`[3]`)
		}
	}
	var act = New("dave",doit)
	var result = make(chan *message.Message)
	go act.Wait(result)
	act.Send(message.New(`[ "hello", "world" ]`))
	act.Send(message.New(`[ "hello", "world" ]`))
	act.Send(message.New(`[ "hello", "world" ]`))
	v,x := (<-result).At(0)
	if v != 3.0 || x != "number" {
		t.Errorf("Wrong number of messages sent")
	}
}

