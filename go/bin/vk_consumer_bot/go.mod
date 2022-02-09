module vk_consumer_bot

go 1.17

require (
	github.com/cornelk/hashmap v1.0.1
	github.com/go-redis/redis/v8 v8.11.4
	github.com/golang/protobuf v1.5.2
	github.com/hashicorp/consul/api v1.12.0
	github.com/nikhovas/diploma/go/lib/proto v0.0.0-00010101000000-000000000000
	github.com/nikhovas/diploma/go/lib/utils v0.0.0-00010101000000-000000000000
	github.com/nikhovas/diploma/go/lib/vk v0.0.0-00010101000000-000000000000
	github.com/rabbitmq/amqp091-go v1.2.0
	google.golang.org/grpc v1.44.0
)

require (
	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dchest/siphash v1.1.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fatih/color v1.9.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/hashicorp/go-hclog v0.12.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.0 // indirect
	github.com/hashicorp/serf v0.9.6 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	golang.org/x/net v0.0.0-20210428140749-89ef3d95e781 // indirect
	golang.org/x/sys v0.0.0-20210423082822-04245dca01da // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace (
	github.com/nikhovas/diploma/go/lib/proto => ../../lib/proto
	github.com/nikhovas/diploma/go/lib/utils => ../../lib/utils
	github.com/nikhovas/diploma/go/lib/vk => ../../lib/vk
)
