package topic

import (
	"testing"
)

func TestNew(t *testing.T) {
	var topic = New("testing")
	if topic.name != "testing" {
		t.Errorf("Invalid name for topic %v", topic.name)
	}
}

func TestSend(t *testing.T) {
	
}

func TestSubscribe(t *testing.T) {

}

func TestUnsubscribe(t *testing.T) {

}
