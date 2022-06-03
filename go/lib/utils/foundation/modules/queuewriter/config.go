package queuewriter

type Config struct {
	Url string `yaml:"url"`
	ExchangeName string `yaml:"exchange-name"`
	QueueName string `yaml:"queue-name"`
}
