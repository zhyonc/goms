package maple

type InvOps uint8

const (
	InvOpsAdd InvOps = iota
	InvOpsUpdateNumber
	InvOpsMove
	InvOpsRemove
	InvOpsItemExp
	InvOpsUpdateBagIndex
	InvOpsUpdateBagNumber
	InvOpsBagRemove
	InvOpsBagToBag
	InvOpsBagNewItem
	InvOpsBagRemoveSlot
)

type InvType int8

const (
	InvTypeEquipped InvType = -1
	InvTypeEquip    InvType = iota
	InvTypeConsume
	InvTypeInstall
	InvTypeEtc
	InvTypeCash
	InvTypeHair
	InvTypeFace
)
