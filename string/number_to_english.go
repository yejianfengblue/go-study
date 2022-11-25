package string

import "fmt"

var NumberToWord = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func main() {

	//fmt.Println(intToWords(8))
	//fmt.Println(intToWords(18))
	//fmt.Println(intToWords(80))
	//fmt.Println(intToWords(88))
	//fmt.Println(intToWords(800))
	//fmt.Println(intToWords(808))
	//fmt.Println(intToWords(810))
	//fmt.Println(intToWords(888))
	//
	//fmt.Println(intToWords(8888))
	//fmt.Println(intToWords(88888))
	//fmt.Println(intToWords(888888))
	fmt.Println(intToWords(8888888))
	fmt.Println(intToWords(88888888))
	fmt.Println(intToWords(888888888))
	fmt.Println(intToWords(8888888888))
	fmt.Println(intToWords(8888888088))

	//var n int
	//fmt.Scan(&n)

}

func intToWords(n int) (words string) {

	billion := n / 1_000_000_000
	n %= 1_000_000_000
	million := n / 1_000_000
	n %= 1_000_000
	thousand := n / 1_000
	n %= 1_000

	if billion > 0 {
		words += xxxToWords(billion) + " billion"
	}
	if million > 0 {
		if billion > 0 {
			words += " "
		}
		words += xxxToWords(million) + " million"
	}
	if thousand > 0 {
		if billion > 0 || million > 0 {
			words += " "
		}
		words += xxxToWords(thousand) + " thousand"
	}
	if n > 0 {
		if billion > 0 || million > 0 || thousand > 0 {
			words += " "
		}
		words += xxxToWords(n)
	}
	return words
}

func xxxToWords(n int) (words string) {

	if n >= 100 {
		hundred := n / 100
		n %= 100
		words += NumberToWord[hundred] + " hundred"
		if n > 0 {
			words += " and "
		}
	}

	if 0 < n && n < 20 {
		words += NumberToWord[n]
	} else {
		r := n % 10
		if r == 0 {
			words += NumberToWord[n]
		} else {
			words += NumberToWord[n-r] + " " + NumberToWord[r]
		}
	}
	return words
}
