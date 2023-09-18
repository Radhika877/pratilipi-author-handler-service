package queue

type MessageStruct struct {
	AuthorId string `json:"authorID"`
}

type QueueStruct struct {
	Message   MessageStruct
	QueueName string
}

func NewQueueStruct(message string, queueName string) QueueStruct {
	return QueueStruct{
		Message:   MessageStruct{message},
		QueueName: queueName,
	}
}
