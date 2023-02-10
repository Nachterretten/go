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
	var romNum []string
	var res int
	var exp2 []string
	fmt.Println("Введите выражение: ")
	fmt.Scan(&str)
	ch := regexp.MustCompile("[+,--,*,/]+")
	z := ch.FindAllString(str, -1)
	var i = [3]byte{'I', 'V', 'X'}
	if str[0] == i[0] || str[0] == i[1] || str[0] == i[2] {
		re2 := regexp.MustCompile("[I,II,III,IV,V,VI,VII,VIII,IX,X]+")
		romNum = re2.FindAllString(str, -1)
		n1 := romanToInt(romNum[0])
		n2 := romanToInt(romNum[1])
		switch z[0] {
		case "+":
			res = n1 + n2
			fmt.Println(intToRoman(res))
		case "-":
			res = n1 - n2
			fmt.Println(intToRoman(res))
		case "/":
			res = n1 / n2
			fmt.Println(intToRoman(res))
		case "*":
			res = n1 * n2
			fmt.Println(intToRoman(res))
		default:
			fmt.Println("Веедите валидное значние")
		}
	} else {

		re := regexp.MustCompile("[0-9]+")
		exp2 = re.FindAllString(str, -1)
		n1, err := strconv.Atoi(exp2[0])
		if err != nil {
			panic("Значение 1 не валидно")
		}
		n2, err := strconv.Atoi(exp2[1])
		if err != nil {
			panic("Значение 2 не валидно")
		}
		switch z[0] {
		case "+":
			res = n1 + n2
			fmt.Println(res)
		case "-":
			res = n1 - n2
			fmt.Println(res)
		case "/":
			res = n1 / n2
			fmt.Println(res)
		case "*":
			res = n1 * n2
			fmt.Println(res)
		}
	}
}
