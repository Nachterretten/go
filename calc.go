package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func romanToInt(s string) int {
	rMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	result := 0
	for k := range s {
		if k < len(s)-1 && rMap[s[k:k+1]] < rMap[s[k+1:k+2]] {
			result -= rMap[s[k:k+1]]
		} else {
			result += rMap[s[k:k+1]]
		}
	}
	return result
}

func intToRoman(num int) string {
	r := [][]string{
		{"", "M", "MM", "MMM"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}}
	n := []int{1000, 100, 10, 1}
	result := ""
	for k, v := range n {
		result += r[k][num/v]
		num = num % v
	}
	return result
}

func main() {

	var str string
	var z []string
	var res int
	var exp2 []string
	var romanMap = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
	}
	var ch int
	fmt.Println("Введите выражение: ")
	fmt.Scan(&str)
	exp, _ := regexp.Compile("[IVX]+")
	tokens := exp.FindAllString(str, -1)
	fmt.Println("Введите выражение:\n1- если вы ввели римские числа.\n2- если вы ввели арабские числа")
	fmt.Scan(&ch)
	switch ch {
	case 1:
		for i := range romanMap {
			if i == tokens[0] || i == tokens[1] {
				ch := regexp.MustCompile("[+,--,*,/]+")
				z = ch.FindAllString(str, -1)
				switch z[0] {
				case "+":
					res = romanToInt(tokens[0]) + romanToInt(tokens[1])
				case "-":
					res = romanToInt(tokens[0]) - romanToInt(tokens[1])
				case "*":
					res = romanToInt(tokens[0]) * romanToInt(tokens[1])
				case "/":
					res = romanToInt(tokens[0]) / romanToInt(tokens[1])
				}
			} else {
				fmt.Println("Введите валидные значния")
			}
		}
		fmt.Println(intToRoman(res))
	case 2:
		ch := regexp.MustCompile("[+,--,*,/]+")
		z = ch.FindAllString(str, -1)
		re := regexp.MustCompile("[0-9]+")
		exp2 = re.FindAllString(str, -1)
		n1, err := strconv.Atoi(exp2[0])
		if err != nil {
			panic("Введите валидное значение")
		}
		n2, err := strconv.Atoi(exp2[1])
		if err != nil {
			panic("Введите валидное значение")
		}
		switch z[0] {
		case "+":
			fmt.Println(n1 + n2)
		case "-":
			fmt.Println(n1 - n2)
		case "*":
			fmt.Println(n1 + n2)
		case "/":
			fmt.Println(n1 + n2)
		}
	}
}
