package types

import "state_machine_executor/state_machine/storageInterfaces"

type TypeInterface interface {
	Set(value interface{}, old interface{}, storage storageInterfaces.IStorage, path string)
	Get(storage storageInterfaces.IStorage, path string) interface{}
}
