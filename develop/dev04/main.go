package main

import (
	"fmt"
	"sort"
	"strings"
)

func isUnique(words *[]string) *map[string]string{
	unique := make(map[string]string)

	for i := range *words {
		word := strings.ToLower((*words)[i])
		if _, ok := unique[word]; !ok {
			unique[word] = sortLetters(&word)
		}
	}
	return &unique
}

func sortLetters(word *string) string {
	rUnique := []rune(*word)
	sort.Slice(rUnique, func(i, j int) bool { return rUnique[i] < rUnique[j]})
	return string(rUnique)
}

func findAnorgam(words *[]string) *map[string][]string {
	sets := make(map[string][]string)
	// сначала удаляю повторяющиеся слова и возвращаю мапу где ключ это уникальное слово, 
	// а значние это отсортированные буквы этого слова
	unique := isUnique(words)
	// создаю множества, где ключ это отсортированные наборы буквы слов, а значние массив строк 
	// со словами, буквенные наборы которых похожи
	for key, value := range *unique {
		sets[value] = append(sets[value], key)
	}
	// в результурующую мапу в значение кладу первое слово из массива отсортированного 
	// множетсва анограмм 
	result := make(map[string][]string, len(sets))
	for key := range sets {
		sort.Strings(sets[key])
		if len(sets[key]) > 1 {
			result[sets[key][0]] = sets[key]
		}
	}

	return &result
}

func main() {
	array := []string{"пятка", "тяпка", "лук", "пяТак", "ЛистОк", "листок", "Абырвалг", "главрыба", "слиток", "столик"}
	sets := findAnorgam(&array)
	for key, set := range *sets {
		fmt.Println(key,"-", set)
	}
}
