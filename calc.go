package main

import "fmt"

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
	var n1, n2 int
	var r1, r2 string
	fmt.Println("1 - Римские числа\n2 - Арабские числа")
	var choiseNum int
	fmt.Scan(&choiseNum)
	if choiseNum == 1 {
		fmt.Println("Введите первое число: ")
		fmt.Scan(&r1)
		fmt.Println("Введите второе число: ")
		fmt.Scan(&r2)

		var choice2 int = 0
		var result2 int = 0

		fmt.Println("1: +")
		fmt.Println("2: -")
		fmt.Println("3: *")
		fmt.Println("4: /")
		fmt.Print("Выберите действие: ")
		fmt.Scan(&choice2)

		switch choice2 {
		case 1:
			result2 = romanToInt(r1) + romanToInt(r2)

			fmt.Printf("Результат суммирования: %s", intToRoman(result2))
		case 2:
			result2 = romanToInt(r1) - romanToInt(r2)
			fmt.Printf("Результат вычитания: %s", intToRoman(result2))
		case 3:
			result2 = romanToInt(r1) * romanToInt(r2)
			fmt.Printf("Результат умножения: %s", intToRoman(result2))
		case 4:
			result2 = romanToInt(r1) / romanToInt(r2)
			fmt.Printf("Рузельтат деление: %s", intToRoman(result2))
		default:
			fmt.Println("не валидные значения")
		}

	} else if choiseNum == 2 {
		fmt.Println("Введите первое число: ")
		fmt.Scan(&n1)
		fmt.Println("Введите второе число: ")
		fmt.Scan(&n2)

		var choice int = 0
		var result int = 0

		fmt.Println("1: +")
		fmt.Println("2: -")
		fmt.Println("3: *")
		fmt.Println("4: /")
		fmt.Print("Выберите действие: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			result = n1 + n2
			fmt.Printf("Результат суммирования: %d", result)
		case 2:
			result = n1 - n2
			fmt.Printf("Результат вычитания: %d", result)
		case 3:
			result = n1 * n2
			fmt.Printf("Результат умножения: %d", result)
		case 4:
			result = n1 / n2
			fmt.Printf("Рузельтат деление: %d", result)
		default:
			fmt.Println("не валидные значения")
		}
	}
}
