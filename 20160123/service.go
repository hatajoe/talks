package itemservice

import (
	"errors"
	"types/itemreceiver"
)

func RegisterItem(m itemreceiver.Receiver, i *Item) error {
	if m.CanReceiveItem(i) {
		return errors.New("err: can't receive item")
	}
	return m.ReceiveItem(i)
}
