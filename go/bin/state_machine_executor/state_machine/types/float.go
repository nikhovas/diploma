package types

import "state_machine_executor/state_machine/storageInterfaces"

type Float struct {
}

func (t *Float) Set(value interface{}, old interface{}, storage storageInterfaces.IStorage, path string) {

}

func (t *Float) Get(storage storageInterfaces.IStorage, path string) interface{} {
	return nil
}
