package types

import "MsgCombiner/stateMachine/storageInterfaces"

type Dict struct {
	Key   TypeInterface
	Value TypeInterface
}

func (t *Dict) Set(value interface{}, old interface{}, storage storageInterfaces.IStorage, path string) {

}

func (t *Dict) Get(storage storageInterfaces.IStorage, path string) interface{} {
	return nil
}
