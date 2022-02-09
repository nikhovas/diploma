package combined

import (
	"control/interfaces"
	"control/modules/kernel"
)

type Combined struct {
	Kernel    *kernel.Kernel
	Staff     map[string]interfaces.Staff
	Consumers map[string]interfaces.Consumer
}
