package main

type MessageQueue struct {
	Msg []string
}

func (m *MessageQueue) Push(msg string) {
	m.Msg = append(m.Msg, msg)
}

func (m *MessageQueue) Pop() string {
	if len(m.Msg) == 0 {
		return "EMPTY"
	}
	msg := m.Msg[0]
	m.Msg = m.Msg[1:]

	return msg
}

func NewMessageQueue() *MessageQueue {
	return new(MessageQueue)
}
