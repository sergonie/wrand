package wrand

import "sort"

type Item struct {
	Value  int
	Weight int
}

type ItemsCollection struct {
	Items []Item
}

func NewItemsCollection(items []Item) *ItemsCollection {
	return &ItemsCollection{Items: items}
}

//@todo: need to check for uniqueness `val`
func (ic *ItemsCollection) Add(val int, weight int) *ItemsCollection {
	ic.Items = append(ic.Items, Item{Value: val, Weight: weight})

	return ic
}

func (ic *ItemsCollection) GetWeightSum() int {
	sum := 0
	for _, item := range ic.Items {
		sum += item.Weight
	}

	return sum
}

func (ic *ItemsCollection) SortByWeightDesc() *ItemsCollection {
	//sort items by weight
	sort.Slice(ic.Items, func(i, j int) bool {
		return ic.Items[i].Weight < ic.Items[j].Weight
	})

	return ic
}

func (ic *ItemsCollection) Count() int {
	return len(ic.Items)
}

func (ic *ItemsCollection) GetAll() []Item {
	return ic.Items
}
