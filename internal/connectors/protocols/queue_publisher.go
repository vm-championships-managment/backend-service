package connectors_protocols

type QueuePublisher interface {
	Send(msg string)
}
