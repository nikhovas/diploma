package types

import "state_machine_executor/state_machine/storageInterfaces"

type Boolean struct {
}

func (t *Boolean) Set(value interface{}, old interface{}, storage storageInterfaces.IStorage, path string) {

}

func (t *Boolean) Get(storage storageInterfaces.IStorage, path string) interface{} {
	return nil
}
