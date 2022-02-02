package types

import "MsgCombiner/stateMachine/storageInterfaces"

type TypeInterface interface {
	Set(value interface{}, old interface{}, storage storageInterfaces.IStorage, path string)
	Get(storage storageInterfaces.IStorage, path string) interface{}
}
