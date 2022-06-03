module telegrammanagerbot

go 1.17

replace (
	github.com/nikhovas/diploma/go/lib/proto => ../../lib/proto
	github.com/nikhovas/diploma/go/lib/utils => ../../lib/utils
)

require (
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
	github.com/golang/protobuf v1.5.2
	github.com/nikhovas/diploma/go/lib/proto v0.0.0-00010101000000-000000000000
	github.com/nikhovas/diploma/go/lib/utils v0.0.0-00010101000000-000000000000
)

require (
	github.com/google/uuid v1.1.2 // indirect
	github.com/rabbitmq/amqp091-go v1.3.0 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)
