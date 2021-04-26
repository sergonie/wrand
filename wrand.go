package wrand

import "sort"

const maxRandNum = 10000

type Picker struct {
	Randomizer Randomizer
}

func NewPicker() *Picker {
	return &Picker{
		Randomizer: MathRandomizer{},
	}
}

func (p *Picker) SetRandomizer(r Randomizer) *Picker {
	p.Randomizer = r
	return p
}

func getProbabilityMap(ic *ItemsCollection) []int {
	weightSum := ic.GetWeightSum()
	probs := make([]int, ic.Count())
	sumProbs := 0

	for k, item := range ic.SortByWeightDesc().Items {
		sumProbs += int(float32(item.Weight) / float32(weightSum) * maxRandNum)
		probs[k] = sumProbs
	}

	return probs
}

func (p *Picker) Pick(ic *ItemsCollection) Item {
	probs := getProbabilityMap(ic)
	count := len(probs)
	rnd := p.Randomizer.Intn(maxRandNum)

	k := sort.Search(count, func(i int) bool {
		return rnd < probs[i]
	})

	if k < count {
		return ic.Items[k]
	}

	return ic.Items[count-1]
}

func (p *Picker) CountPicksByValues(ic *ItemsCollection, picks int) map[int]int {
	res := make(map[int]int, ic.Count()+1)

	for i := 0; i < picks; i++ {
		item := p.Pick(ic)
		res[item.Value]++
	}

	return res
}
