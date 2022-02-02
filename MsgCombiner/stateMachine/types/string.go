package types

import "MsgCombiner/stateMachine/storageInterfaces"

type String struct {
}

func (t *String) Set(value interface{}, old interface{}, storage storageInterfaces.IStorage, path string) {

}

func (t *String) Get(storage storageInterfaces.IStorage, path string) interface{} {
	return nil
}
