package actions

//type ManagerBot struct {
//	GenericAction
//}
//
//func NewManagerBot(genericAction *GenericAction) ActionInterface {
//	return &ManagerBot{GenericAction: *genericAction}
//}
//
//func (a *ManagerBot) Run(storage *localStorage.Storage) {
//	ManagerBotFunc(storage, a.Arguments, a.Return)
//}
//
//func ManagerBotFunc(storage *localStorage.Storage, Arguments Arguments, Return Returns) {
//	botService := storage.KvStorage.Get("botService").(string)
//	groupId, _ := strconv.Atoi(storage.KvStorage.Get("groupId").(string))
//	userId, _ := strconv.Atoi(storage.KvStorage.Get("userId").(string))
//
//	msg := storage.KvStorage.Get("message")
//	storage.MessageDeque.PushFront(msg)
//}
