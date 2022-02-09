package kernel

import (
	"github.com/nikhovas/diploma/go/lib/proto/consumer_bot"
	qw "github.com/nikhovas/diploma/go/lib/proto/question_worker"
	"github.com/nikhovas/diploma/go/lib/utils/distfs"
)

type Kernel struct {
	DistFsRoot        *distfs.Root
	QwClient          qw.QuestionWorkerClient
	ConsumerBotClient consumer_bot.VkServerClient
}
