package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func isValid(str []rune) bool {
	if str[0] >= '0' && str[0] <= '9' {
		return false
	}
	return true
}

func repitSign(sign rune, counter int) string {
	var str string

	for ; counter > 0; counter-- {
		str += string(sign)
	}
	return str
}

func checkSignAfterEscape(sign rune) rune {
	switch sign {
	case 'n':
		return '\n'
	case 'b':
		return '\b'
	case 'r':
		return '\r'
	case 't':
		return '\t'
	case 'v':
		return '\v'
	default:
		return sign
	}
}

func stringHandler(str string) (string, error) {
	var (
		sign    rune
		counter int
		newStr  string
	)
	if len(str) == 0 {
		return "", nil
	}
	rStr := []rune(str)
	if isValid(rStr) {
		for i := 0; i < len(rStr); i++ {
			if rStr[i] == '\\' {
				i++
				sign = checkSignAfterEscape(rStr[i])
				newStr += string(sign)
			} else if rStr[i] >= '0' && rStr[i] <= '9' {
				counter, _ = strconv.Atoi(string(rStr[i]))
				newStr += repitSign(sign, counter-1)
			} else {
				sign = rStr[i]
				newStr += string(sign)
			}
		}
		return newStr, nil
	}
	return "", fmt.Errorf("%s: передана некорректная строка", "stringHandler")
}

func main() {
	result, err := stringHandler(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(result)
	os.Exit(0)
}
