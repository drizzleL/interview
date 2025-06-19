package main

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	supToRecipe := map[string][]int{}
	needs := make([]int, len(recipes))
	for i := range recipes {
		needs[i] = len(ingredients[i])
		for _, ind := range ingredients[i] {
			supToRecipe[ind] = append(supToRecipe[ind], i)
		}
	}
	var q []int
	for _, sup := range supplies {
		for _, rec := range supToRecipe[sup] {
			needs[rec] -= 1
			if needs[rec] == 0 {
				q = append(q, rec)
			}
		}
	}
	var ret []string
	for len(q) != 0 {
		pop := q[len(q)-1]
		q = q[:len(q)-1]
		ret = append(ret, recipes[pop])
		for _, rec := range supToRecipe[recipes[pop]] {
			needs[rec] -= 1
			if needs[rec] == 0 {
				q = append(q, rec)
			}
		}
	}
	return ret
}
