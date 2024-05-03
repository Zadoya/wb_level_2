package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения							//done
-B - "before" печатать +N строк до совпадения							//done
-C - "context" (A+B) печатать ±N строк вокруг совпадения				//done
-c - "count" (количество строк)											//done
-i - "ignore-case" (игнорировать регистр)								//done
-v - "invert" (вместо совпадения, исключать)							//done
-F - "fixed", точное совпадение со строкой, не паттерн					//done
-n - "line num", печатать номер строки                                  //done

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

type Flags struct {
	after       *int
	before      *int
	context     *int
	count       *bool
	ignoreCase	*bool
	invert      *bool
	fixed       *bool
	lineNum    	*bool
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
		log.Fatal(err)
	}

	return &lines
}

func numeric(src *[]string, start, end, founded int) {
	for i := start; i <= end; i++ {
		if i != founded {
			(*src)[i] = fmt.Sprintf("%d-%s", i + 1, (*src)[i])
		} else {
			(*src)[i] = fmt.Sprintf("%d:%s", i + 1, (*src)[i])
		}
	}
}

func returnLines(src *[]string, founded *[]int, flags Flags) *[]string {
	var start, end int
	
	dest := make([]string, 0, len(*founded))
	for i, idx := range *founded {
		if i > 0 {
			start = idx - min(max(*flags.before, *flags.context), idx - (*founded)[i - 1] - 1)
		} else {
			start = idx - max(*flags.before, *flags.context)
		}
		if start < 0 {
				start = 0
		}
		
		if i + 1 < len(*founded) {
			end = idx + min(max(*flags.after, *flags.context), (*founded)[i + 1] - idx - 1)
		} else {
			end = idx + max(*flags.after, *flags.context)
		}
		if end > len(*src) {
			end = len(*src)
		}

		if len(dest) > 0 {
			if *flags.context > 0 {
				dest = append(dest, "--\n--")
			} else if *flags.after > 0 || *flags.before > 0 {
				dest = append(dest, "--")
			}
		}
		if *flags.lineNum {
			numeric(src, start, end, idx)
		}
		dest = append(dest, (*src)[start:end + 1]...)
	}
	
	return &dest
}

func grep(data *[]string, pattern string, flags Flags) (*[]string, int) {
	if pattern == "" {
		return data, len(*data)
	}

	contains := make([]int, 0, len(*data))
	counter  := 0
	for i := range *data {
		if *flags.ignoreCase {
			(*data)[i] = strings.ToLower((*data)[i])
			pattern = strings.ToLower(pattern)
		}
		
		if *flags.fixed && ((strings.Compare((*data)[i], pattern) == 0) != *flags.invert) {
			contains = append(contains, i)
			counter++
		} else if (strings.Contains((*data)[i], pattern) != *flags.invert) {
			contains = append(contains, i)
			counter++
		}
	}
	return returnLines(data, &contains, flags), counter
}

func main() {
	flags := Flags{
		after: flag.Int("A", 0, "Печатать +N строк после совпадения"),
		before: flag.Int("B", 0, "Печатать +N строк до совпадения"),
		context: flag.Int("C", 0, "Печатать ±N строк вокруг совпадения"),
		count: flag.Bool("c", false, "Кол-во совпадающих строк"),
		ignoreCase: flag.Bool("i", false, "Bгнорировать регистр"),
		invert: flag.Bool("v", false, "Вместо совпадения, исключать"),
		fixed: flag.Bool("F", false, "Точное совпадение со строкой"),
		lineNum: flag.Bool("n", false, "Напечатать номер строки"),
	}

	flag.Parse()

	var pattern, file string

	switch len(flag.Args()) {
	case 0:
		fmt.Println("необходимо указать паттерн и имя файла")
		os.Exit(1)
	case 1:
		fmt.Println("необходимо указать путь до файла")
	case 2:
		pattern = flag.Args()[0]
		file = flag.Args()[1]
	default:
		fmt.Println("слишком много аргументов")
		os.Exit(1)
	}

	dataFile := readfile(file)

	dataFile, counter := grep(dataFile, pattern, flags)

	if *flags.count {
		fmt.Println(counter)
	} else {
		for _, line := range *dataFile {
			fmt.Println(line)
		}
	}
}
