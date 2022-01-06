package localStorage

import "MsgCombiner/stateMachine/storageInterfaces"

type DataElement struct {
	Memory  string
	Type    string
	Default interface{}
}

type KvStorageElement struct {
	Memory       string
	Type         string
	Modification bool
	Value        interface{}
	Default      interface{}
}

type KvStorage struct {
	storage map[string]*KvStorageElement
	helper  storageInterfaces.IGetter
}

func (kv *KvStorage) Init(descriptions map[string]DataElement, helper storageInterfaces.IGetter) {
	kv.storage = make(map[string]*KvStorageElement, 4)
	for key, value := range descriptions {
		kv.storage[key] = &KvStorageElement{
			Memory:       value.Memory,
			Type:         value.Type,
			Modification: false,
			Value:        nil,
			Default:      value.Default,
		}
	}

	kv.helper = helper
}

func (kv *KvStorage) Set(key string, value interface{}) {
	current := kv.storage[key]
	if current.Type == "stringList" {
		v1 := value.([]string)
		v2 := current.Value.([]string)
		if len(v1) == len(v2) {
			equals := true
			for i := 0; i < len(v1); i++ {
				if v1[i] != v2[i] {
					equals = false
					break
				}
			}
			if equals {
				return
			}
		}
	} else {
		if current.Value == value {
			return
		}
	}

	current.Modification = true
	current.Value = value
}

func (kv *KvStorage) Get(key string) interface{} {
	v := kv.storage[key]
	if v.Memory == "const" {
		return v.Default
	}

	if v.Value != nil {
		return v.Value
	}

	v.Value = kv.helper.Get(key, v.Type)
	if v.Value != nil {
		return v.Value
	}

	v.Value = v.Default
	if v.Value == nil {
		if v.Type == "stringList" {
			v.Value = make([]string, 0)
		}
	}
	return v.Value
}

func (kv *KvStorage) GetUpdatedMemory(memoryType string) map[string]interface{} {
	updated := make(map[string]interface{}, 4)
	for key, value := range kv.storage {
		if value.Modification && value.Memory == memoryType {
			updated[key] = value.Value
		}
	}

	return updated
}

func (kv *KvStorage) GetUpdatedLongMemory() map[string]interface{} {
	return kv.GetUpdatedMemory("long")
}

func (kv *KvStorage) GetUpdatedHardMemory() map[string]interface{} {
	return kv.GetUpdatedMemory("hard")
}
