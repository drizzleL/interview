package main

type WordsFrequency struct {
	dict map[string]int
}

func ConstructorW(book []string) WordsFrequency {
	dict := map[string]int{}
	for _, w := range book {
		dict[w]++
	}
	return WordsFrequency{
		dict: dict,
	}

}

func (this *WordsFrequency) Get(word string) int {
	return this.dict[word]
}
