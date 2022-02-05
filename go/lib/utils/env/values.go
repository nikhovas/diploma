package env

func GetControllerGrpcHost() string {
	return GetEnv("CONTROLLER_GRPC_HOST", "localhost:7777")
}

func GetVkConsumerBotGrpcHost() string {
	return GetEnv("VK_CONSUMER_BOT_GRPC_HOST", "localhost:5555")
}

func GetAmqpUrl() string {
	return GetEnv("AMQP_URL", "amqp://guest:guest@localhost:5672/")
}

func GetStateMachineExecutorAmqpQueue() string {
	return GetEnv("STATE_MACHINE_EXECUTOR_AMQP_QUEUE", "action_events")
}
