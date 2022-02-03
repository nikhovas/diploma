package storageInterfaces

type IGetter interface {
	Get(key string, valueType string) interface{}
}

type ISetter interface {
	Set(key string, value interface{})
}

type IStorage interface {
	IGetter
	ISetter
}
