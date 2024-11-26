package config

type Settings struct {
	Port          int           `json:"port"`
	QueryService  string        `json:"queryService"`
	TripService   string        `json:"tripService"`
	KafkaSettings KafkaSettings `json:"kafka"`
}

type KafkaSettings struct {
	Brokers []string    `json:"brokers"`
	Topics  KafkaTopics `json:"topics"`
}

type KafkaTopics struct {
	InsertTopic string `json:"insert"`
}
