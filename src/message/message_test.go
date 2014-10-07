package message

import (
	"testing"
	"reflect"
)

func TestNew(t *testing.T) {
	var m1 = New([]byte(`[ "hello", "world", 1,2,3 ]`)) 
	if str := m1.String(); `["hello","world",1,2,3]` != str {
		t.Errorf("String not serialized correctly got %v", str )
	}
	var m2 = New([]byte(`[{ "foo" : "bar" }]`))
	if str := m2.String(); `[{"foo":"bar"}]` != str {
		t.Errorf("String not serialized correctly got %v", str )
	}
}

func TestNewWithSender(t *testing.T) {
	var m = NewWithSender([]byte(`[ "hello", "world", 1,2,3 ]`),"dave") 
	if m.sender != "dave" {
		t.Errorf("Incorrect sender %s", m.sender)
	}
}

func TestNewWithTopic(t *testing.T) {
	var m = NewWithTopic([]byte(`[ "hello", "world", 1,2,3 ]`),"kayaking") 
	if m.topic != "kayaking" {
		t.Errorf("Incorrect sender %s", m.sender)
	}
}

func TestNewWithSenderAndTopic(t *testing.T) {
	var m = NewWithSenderAndTopic([]byte(`[ "hello", "world", 1,2,3 ]`),"dave","canals") 
	if m.sender != "dave" {
		t.Errorf("Incorrect sender %s", m.sender)
	}
	if m.topic != "canals" {
		t.Errorf("Incorrect topic %s", m.topic)
	}
}

func TestRecipient(t *testing.T) {
	var m1 = New([]byte(`[ "hello", "world", 1,2,3 ]`)) 
	if recipient := m1.Recipient(); recipient != "hello" {
		t.Errorf("Incorrect recipient %v for %v", recipient,m1)
	}
	var m2 = New([]byte(`[{ "foo" : "bar" }]`))
	if recipient := m2.Recipient(); recipient != "*" {
		t.Errorf("Incorrect recipient %v for %v", recipient, m2)
	}
}

func TestAt(t *testing.T) {
	var i = 0
	var m = New([]byte(`[ "hello", "world", 1, 2.2, {"foo":"bar"}, [1,2,3], false ]`))
	if u,v := m.At(i); u != "hello" || v != "string" {
		t.Errorf("Invalid arg at %v %v %v",i,u,v)
	}
	i += 1
	if u,v := m.At(i); u != "world" || v != "string" {
		t.Errorf("Invalid arg at %v %v %v",i,u,v)
	}
	i += 1
	if u,v := m.At(i); u != 1.0 || v != "number" {
		t.Errorf("Invalid arg at %v %v %v",i,u,v)
	}
	i += 1
	if u,v := m.At(i); u != 2.2 || v != "number" {
		t.Errorf("Invalid arg at %v %v %v",i,u,v)
	}
	i += 1
	if u,v := m.At(i); !reflect.DeepEqual(u,map[string]interface{}{"foo":"bar"}) || v != "object" {
		t.Errorf("Invalid arg at %v %v %v",i,u,v)
	}
	i += 1
	if u,v := m.At(i); u.([]interface{})[0] != 1.0 || u.([]interface{})[1] != 2.0 || u.([]interface{})[2] != 3.0   || v != "array" {
		t.Errorf("Invalid arg at %v %v %v",i,u,v)
	}
	i += 1
	if u,v := m.At(i); u != false  || v != "boolean" {
		t.Errorf("Invalid arg at %v %v %v",i,u,v)
	}
	i += 1
}
