# WRAND

Lightweight golang library for obtaining a weighted random item.

## Installation

```shell
go get -u github.com/sergonie/wrand
```

## Example

```go
package main

import (
	"fmt"
	"github.com/sergonie/wrand"
)

func main() {
	myDataset := []struct {
		Name   string
		Weight int
	}{
		{
			Name:   "Max",
			Weight: 80,
		},
		{
			Name:   "El",
			Weight: 110,
		},
		{
			Name:   "Gabriella",
			Weight: 45,
		},
	}

	items := &wrand.ItemsCollection{}
	for k, v := range myDataset {
		items.Add(k, v.Weight)
	}

	selectedItem := wrand.NewPicker().Pick(items)

	fmt.Printf("Winner: %s\n", myDataset[selectedItem.Value].Name)
}
```
