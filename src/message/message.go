package message

import (
	"encoding/json"
)

type Message struct {
	sender string
	topic string
	value []interface{}
}

func FromBytes(data []byte) *Message {
	var m = &Message{}
	if err := json.Unmarshal(data, &m.value); err != nil {
		return &Message{}
	}
	return m
}

func FromString(s string) *Message {
	return FromBytes([]byte(s))
}

func New(s string) *Message {
	return FromBytes([]byte(s))
}

func NewWithSender(data string, sender string) *Message {
	var m = New(data)
	m.sender = sender
	return m
}

func NewWithTopic(data string, topic string) *Message {
	var m = New(data)
	m.topic = topic
	return m
}

func NewWithSenderAndTopic(data string, sender string, topic string) *Message {
	var m = New(data)
	m.sender = sender
	m.topic = topic
	return m
}

func (m *Message) String() string {
	var value []interface{} = m.value
	if b, err := json.Marshal(value); err == nil {
		return string(b[:])
	}
	return "[]"
}

func (m *Message) Recipient() string {
	switch m.value[0].(type) {
	case string:
		return m.value[0].(string)
	default:
		return "*"
	}
}

func (m *Message) Topic() string {
	return m.topic
}

func (m *Message) At( i int ) (interface{}, string) {
	switch m.value[i].(type) {
		case bool:
			return m.value[i].(bool), "boolean"
		case string:
			return m.value[i].(string), "string"
		case float64:
			return m.value[i].(float64), "number"
		case []interface{}:
			return m.value[i].([]interface{}), "array"
		case map[string]interface{}:
			return m.value[i].(map[string]interface{}), "object"
		default:
			return nil, "undefined"
	}
}

