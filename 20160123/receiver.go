package itemreceiver

import (
	"models/item"
)

type Receiver struct {
	CanReceiveItem(*item.Item) error
	ReceiveItem (*item.Item) error
}

