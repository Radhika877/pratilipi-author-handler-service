package lib

type Config struct {
	Port              string `json:"Port"`
	Name              string `json:"Name"`
	Version           string `json:"Version"`
	VersionDate       string `json:"VersionDate"`
	MongodbURI        string `json:"MongodbURI"`
	MongoDBCollection *MongoDBCollection
	ConfigDocID       string `json:"ConfigDocID"`
	AMQPEndpoint      string `json:"AMQPEndpoint"`
	QueueName         string `json:"QueueName"`
}
