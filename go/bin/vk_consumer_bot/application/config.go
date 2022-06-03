package application

import (
	"github.com/nikhovas/diploma/go/lib/utils/foundation/modules/queuewriter"
	"vk_consumer_bot/modules/grpcserver"
	"vk_consumer_bot/modules/messageobserver"
	"vk_consumer_bot/modules/readmessagestosend"
	"vk_consumer_bot/modules/workingbotsupdater"
)

type Config struct {
	EventsQueueWriter queuewriter.Config `yaml:"events-queue-writer"`
	MessageObserver messageobserver.Config `yaml:"message-observer"`
	GrpcServer grpcserver.Config `yaml:"grpc-server"`
	WorkingBotsUpdater workingbotsupdater.Config `yaml:"working-bots-updater"`
	ReadMessagesToSend readmessagestosend.Config `yaml:"read-messages-to-send"`
}
