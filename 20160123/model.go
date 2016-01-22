package model

import (
	"models/item"
)

type Member struct {
	ID    int64
	Items []item.Item
}

func (m *Member) RegisterItem(i item.Item) {
	m.Items = append(m.Items, i)
}
