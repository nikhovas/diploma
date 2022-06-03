package common

import "github.com/nikhovas/diploma/go/lib/utils/distvars"

type MetaShopId struct {
	DirInfo distvars.MetaDir
}

func (dv *MetaShopId) CdQuestionSources() *distvars.RedisDictStringString {
	return dv.DirInfo.GetRedisDictStringStringChild("question-sources")
}

func (dv *MetaShopId) CdQuestions() *distvars.RedisStringSet {
	return dv.DirInfo.GetRedisStringSetChild("questions")
}

func (dv *MetaShopId) CdQa() *distvars.RedisDictStringString {
	return dv.DirInfo.GetRedisDictStringStringChild("qa")
}

func (dv *MetaShopId) CdShopItems() *distvars.RedisString {
	return dv.DirInfo.GetRedisStringChild("shop-items")
}
