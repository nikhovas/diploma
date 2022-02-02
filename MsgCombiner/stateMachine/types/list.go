package types

import "MsgCombiner/stateMachine/storageInterfaces"

type List struct {
	Element TypeInterface
}

func (t *List) Set(value interface{}, old interface{}, storage storageInterfaces.IStorage, path string) {

}

func (t *List) Get(storage storageInterfaces.IStorage, path string) interface{} {
	return nil
}
