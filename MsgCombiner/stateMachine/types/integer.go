package types

import "MsgCombiner/stateMachine/storageInterfaces"

type Integer struct {
}

func (t *Integer) Set(value interface{}, old interface{}, storage storageInterfaces.IStorage, path string) {

}

func (t *Integer) Get(storage storageInterfaces.IStorage, path string) interface{} {
	return nil
}
