package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

const maxResultCount = 10

// Регулярка для разделителя слов в задании со звёздочкой: любой символ, кроме буквы, цифры и дефиса
var regexSplitter = regexp.MustCompile("[^\\p{L}\\d-]+")

// Top10 без звёздочки
func Top10(str string) []string {
	// Срезаем углы: нет смысла проходить по всему методу, если исходная строка пустая
	if str == "" {
		return []string{}
	}

	// Вытаскиваем слова из строки по пробельным разделителям
	words := strings.Fields(str)

	// Формируем словарь: ключ - слово, значение - кол-во повторов этого слова
	countsMap := countWords(words)

	// Сортируем словарь и отдаём топ10 слов
	return top10FromCountsMap(countsMap)
}

// Top10WithAsterisk со звёздочкой
func Top10WithAsterisk(str string) []string {
	// Срезаем углы: нет смысла проходить по всему методу, если строка пустая
	if str == "" {
		return []string{}
	}

	// Для разделения слов используем регулярку
	words := regexSplitter.Split(str, -1)

	// Формируем словарь с приведением слов к нижнему регистру и удалением тире из списка слов
	countsMap := countWordsWithAsterisk(words)

	return top10FromCountsMap(countsMap)
}

// Функция формирует словарь из массива слов
// В ключи записываются слова, в значения - кол-во их повторений
func countWords(words []string) map[string]int {
	result := map[string]int{}
	for _, word := range words {
		if word == "" {
			continue
		}
		result[word]++
	}

	return result
}

// Функция формирует словарь из массива слов
// В ключи записываются слова, в значения - кол-во их повторений
// Слова приводятся к нижнему регистру, игнорируются тире в качестве слов
func countWordsWithAsterisk(words []string) map[string]int {
	result := map[string]int{}
	for _, word := range words {
		if word == "-" || word == "" {
			continue
		}
		result[strings.ToLower(word)]++
	}

	return result
}

// Функция работает со словарём, содержащим слова и кол-во их повторений
// Возвращает топ10 самых часто повторяемых слов
// При равном числе повторений слова сортируются лексикографически (т.е по алфавиту )))
func top10FromCountsMap(countsMap map[string]int) []string {
	type wcStruct struct {
		word  string
		count int
	}

	// Можем точно рассчитать длину слайса, поэтому обойдёмся без аппенда и реаллокаций
	wc := make([]wcStruct, len(countsMap))
	i := 0
	for k, v := range countsMap {
		wc[i] = wcStruct{k, v}
		i++
	}

	// Сортируем слайс. Если частота повторения слов одинакова, то сортируем лексикографически по возрастанию
	// иначе по частоте по убыванию
	sort.Slice(wc, func(i, j int) bool {
		if wc[i].count == wc[j].count {
			return wc[i].word < wc[j].word
		}

		return wc[i].count > wc[j].count
	})

	// Можем вычислить длину результирующего слайса и обойтись без аппенда и реаллокаций
	maxResCount := maxResultCount
	if len(wc) < maxResultCount {
		maxResCount = len(wc)
	}

	result := make([]string, maxResCount)
	for i := 0; i < maxResCount; i++ {
		result[i] = wc[i].word
	}

	return result
}
