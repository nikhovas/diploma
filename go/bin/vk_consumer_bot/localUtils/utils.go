package localUtils

import ctrl "github.com/nikhovas/diploma/go/lib/proto/controller"

func NewVkGroupIdShopKey(groupId int) *ctrl.ShopKey {
	return &ctrl.ShopKey{
		Key: &ctrl.ShopKey_VkConsumer{
			VkConsumer: &ctrl.VkConsumerShopKey{
				Key: &ctrl.VkConsumerShopKey_GroupId{
					GroupId: int64(groupId),
				},
			},
		},
	}
}
