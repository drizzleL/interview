package main

import "container/heap"

type Food struct {
	Rating  int
	Name    string
	Idx     int
	Cuisine string
}

type Cuisine []*Food

func (h Cuisine) Len() int { return len(h) }
func (h Cuisine) Less(i, j int) bool {
	a, b := h[i], h[j]
	if a.Rating == b.Rating {
		return a.Name < b.Name
	}
	return a.Rating > b.Rating
}
func (h Cuisine) Swap(i, j int) {
	a, b := h[i], h[j]
	h[i], h[j] = h[j], h[i]
	a.Idx, b.Idx = b.Idx, a.Idx
}

func (h *Cuisine) Push(x any) {
	*h = append(*h, x.(*Food))
}
func (h *Cuisine) Top() *Food {
	return (*h)[0]
}

func (h *Cuisine) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type FoodRatings struct {
	dict map[string]*Cuisine
	food map[string]*Food
}

func FoodConstructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	dict := map[string]*Cuisine{}
	foodDict := map[string]*Food{}
	for i := range foods {
		f := foods[i]
		c := cuisines[i]
		r := ratings[i]
		if dict[c] == nil {
			dict[c] = &Cuisine{}
		}
		food := &Food{
			Rating:  r,
			Name:    f,
			Cuisine: c,
			Idx:     dict[c].Len(),
		}
		dict[c].Push(food)
		foodDict[f] = food
	}
	for _, h := range dict {
		heap.Init(h)
	}
	return FoodRatings{
		dict: dict,
		food: foodDict,
	}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	f := this.food[food]
	h := this.dict[f.Cuisine]
	f.Rating = newRating
	heap.Fix(h, f.Idx)
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	h := this.dict[cuisine]
	return h.Top().Name
}
