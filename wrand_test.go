package wrand

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

const maxRndPercentsError = 5

type testPickerPickCase struct {
	Name  string
	Items *ItemsCollection
	Picks int
}

func (t *testPickerPickCase) hasValue(value int) bool {
	for _, i := range t.Items.GetAll() {
		if i.Value == value {
			return true
		}
	}

	return false
}

func (t *testPickerPickCase) calcPercents() map[int]float64 {
	//weight sum
	weightSum := 0
	for _, v := range t.Items.GetAll() {
		weightSum += v.Weight
	}

	//percents
	res := make(map[int]float64, t.Items.Count())
	for _, v := range t.Items.GetAll() {
		res[v.Value] = float64(v.Weight) / float64(weightSum) * 100
	}

	return res
}

func TestPicker_Pick(t *testing.T) {
	var cases = []testPickerPickCase{
		{
			Name: "equal weight",
			Items: NewItemsCollection(
				[]Item{
					{
						Value:  1,
						Weight: 50,
					},
					{
						Value:  2,
						Weight: 50,
					},
					{
						Value:  3,
						Weight: 50,
					},
				}),
			Picks: 1000,
		},
		{
			Name: "weights with a large spread",
			Items: NewItemsCollection(
				[]Item{
					{
						Value:  1,
						Weight: 50,
					},
					{
						Value:  2,
						Weight: 100,
					},
					{
						Value:  3,
						Weight: 5000,
					},
				}),
			Picks: 1000,
		},
		{
			Name: "weights with a small spread",
			Items: NewItemsCollection(
				[]Item{
					{
						Value:  1,
						Weight: 50,
					},
					{
						Value:  2,
						Weight: 55,
					},
					{
						Value:  3,
						Weight: 45,
					},
				}),
			Picks: 1000,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			vPicks := map[int]int{}
			vPercents := c.calcPercents()

			for i := 0; i < c.Picks; i++ {
				r := NewPicker().Pick(c.Items)
				assert.True(t, c.hasValue(r.Value))

				vPicks[r.Value]++
			}

			//calc random percents by value
			for _, v := range c.Items.GetAll() {
				vPercent := float64(vPicks[v.Value]) / float64(c.Picks) * 100

				t.Log(fmt.Sprintf(
					"Expected random percents: %.2f±%d; Actual: %.2f",
					vPercents[v.Value], maxRndPercentsError, vPercent,
				))
				assert.True(t, math.Abs(vPercent-vPercents[v.Value]) <= maxRndPercentsError)
			}

		})
	}

}
