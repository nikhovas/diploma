package readmessagestosend

import "github.com/nikhovas/diploma/go/lib/utils/foundation/modules/queuereader"

type Config struct {
	QueueReader queuereader.Config `yaml:"queue-reader"`
}
