package main

/*
	=== Утилита cut ===

	Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

	Поддержать флаги:
	-f - "fields" - выбрать поля (колонки)
	-d - "delimiter" - использовать другой разделитель
	-s - "separated" - только строки с разделителем

	Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fields := flag.String("f", "", "выбрать поля (колонки)")
	delimiter := flag.String("d", "\t", "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")

	flag.Parse()

	numbersOfFields := make([]int, 0, 2) 
	if *fields == ""{
		log.Fatal("list: list value can not be empty")
	} else {
		tmp := strings.Split(*fields, "-")
		if len(tmp) > 2 {
			log.Fatal("list: illegal list value")
		}
		if len(tmp) == 2 && (tmp[0] == "" || tmp[1] == ""){
			for i := 0; i < len(tmp); i++ {
				number, _ := strconv.Atoi(tmp[i])
				numbersOfFields = append(numbersOfFields, number)
			}
		} else {
			for i := 0; i < len(tmp); i++ {
				number, err := strconv.Atoi(tmp[i])
				if err != nil {
					log.Fatal("list: illegal list value")
				}
				if number == 0 {
					log.Fatal("list: values may not include zero")
				}
				numbersOfFields = append(numbersOfFields, number)
			}
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		words := strings.Split(str, *delimiter)

		if !(*separated && len(words) == 1) {
			if len(words) < numbersOfFields[0] {
				fmt.Println()
			} else {
				if numbersOfFields[0] != 0 && numbersOfFields[1] != 0 {
					if numbersOfFields[1] < len(words) {
						fmt.Println(strings.Join(words[numbersOfFields[0] - 1:numbersOfFields[1]], *delimiter))
					} else {
						fmt.Println(strings.Join(words[numbersOfFields[0] - 1:], *delimiter))
					}
				} else if numbersOfFields[0] == 0 {
					if numbersOfFields[1] < len(words) {
						fmt.Println(strings.Join(words[:numbersOfFields[1]], *delimiter))
					} else {
						fmt.Println(str)
					}
				} else if numbersOfFields[1] == 0 {
					fmt.Println(strings.Join(words[numbersOfFields[0] - 1:], *delimiter))
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}