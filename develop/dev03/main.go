package main

/*
=== Утилита sort ===
Отсортировать строки в файле по аналогии
с консольной утилитой sort (man sort — смотрим описание и основные параметры): на входе подается файл c несортированными строками, на выходе — файл с отсортированными.
# Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительное

# Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)


func myAtoi(s string) int {

	var output int

	sign := false
	if string(s[0]) == "-" {
		sign = true
		s = s[1:]
	} else if string(s[0]) == "+" {
		s = s[1:]
	}

	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		tempNum := int(s[i]) - '0'
		output = output * 10 + tempNum
	}

	if sign {
		output = output * -1
	}

	return output
}

func readfile(path string) *[]string {

	var lines []string

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	
	return &lines
}

func recordFile(path string, data *[]string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for i := range *data {
		file.WriteString((*data)[i] + "\n")
	}
}

func recordCmd(data *[]string) {
	for i := range *data {
		fmt.Println((*data)[i])
	}
}

func uniqueString(data *[]string, n bool) {
	if n {
		for i := 1; i < len(*data); i++ {
			corrent := myAtoi(strings.ReplaceAll((strings.TrimSpace((*data)[i])), " ", ""))
			prev := myAtoi(strings.ReplaceAll((strings.TrimSpace((*data)[i - 1])), " ", ""))
			if i > 0 && corrent == prev {
				(*data) = append((*data)[:i], (*data)[i+1:]...)
				i--
			}
		}
	} else {
		for i := 0; i < len(*data); i++ {
			if i > 0 && strings.Compare((*data)[i], (*data)[i-1]) == 0 {
				(*data) = append((*data)[:i], (*data)[i+1:]...)
				i--
			}
		}
	}

}

func reverse(data *[]string) {
	n := len(*data)
	for i := 0; i < n/2; i++ {
		(*data)[i], (*data)[n-i-1] = (*data)[n-i-1], (*data)[i]
	}
}

func sortNums(data *[]string) {
	nums := make([]int, 0, len(*data))
	for i := range (*data) {
		trimed := strings.ReplaceAll((strings.TrimSpace((*data)[i])), " ", "")
		num := myAtoi(trimed)
		nums = append(nums, num)
		tmp := (*data)[i]
		j := i - 1
		for ; j >= 0 && nums[j] > num; {
			nums[j+1] = nums[j]
			(*data)[j+1] = (*data)[j]
			j--
		}
		nums[j+1] = num
		(*data)[j+1] = tmp
	}
}

func partition(data *[]string, low, high, column int) int {
	pivot := (*data)[high]
	i := low
	for j := low; j < high; j++ {
		pivot := string([]byte(pivot)[column:])
		copyData := string([]byte((*data)[j])[column:])
		if strings.Compare(copyData, pivot) < 0 {
			(*data)[i], (*data)[j] = (*data)[j], (*data)[i]
			i++
		}
	}
	(*data)[i], (*data)[high] = (*data)[high], (*data)[i]
	return i
}

func quickSort(data *[]string, low, high, column int) {
	if low < high {
		p := partition(data, low, high, column)
		quickSort(data, low, p-1, column)
		quickSort(data, p+1, high, column)
	}
}

func sortByColumn(data *[]string, column int) {
	quickSort(data, 0, len(*data)-1, column)
}

func main() {
	k := flag.Int("k", -1, "указание колонки для сортировки")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "сортировать в обратном порядке")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")
	o := flag.String("o", "sorted_", "указание результирующего файла файла")

	flag.Parse()

	var file string

	switch len(flag.Args()) {
	case 0:
		fmt.Println("необходимо указать имя файла, который нужно отсортировать")
		os.Exit(1)
	case 1:
		file = flag.Args()[0]
	default:
		fmt.Println("необходимо указать путь до одного файла")
		os.Exit(1)
	}

	dataFile := readfile(file)

	if *n {
		sortNums(dataFile)
	} else if *k > -1 {
		sortByColumn(dataFile, *k)
	} else {
		quickSort(dataFile, 0, len(*dataFile) - 1, 0)
	}
	if *u {
		uniqueString(dataFile, *n)
	}
	if *r {
		reverse(dataFile)
	}
	if *o == "sorted_" {
		*o += file
	}

	recordFile(*o, dataFile)
	recordCmd(dataFile)
}
