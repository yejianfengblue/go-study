package string

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

var numToWord = map[int]string{

	1:  "壹",
	2:  "贰",
	3:  "叁",
	4:  "肆",
	5:  "伍",
	6:  "陆",
	7:  "柒",
	8:  "捌",
	9:  "玖",
	10: "拾",
}

func TestRmb(t *testing.T) {
	ass := assert.New(t)
	ass.Equal("壹元整", convert(1))
	ass.Equal("壹元伍角", convert(1.5))
	ass.Equal("壹元伍角伍分", convert(1.55))
	ass.Equal("伍角伍分", convert(0.55))
	ass.Equal("拾元整", convert(10))
	ass.Equal("拾壹元整", convert(11))
	ass.Equal("拾壹元伍角", convert(11.5))
	ass.Equal("壹佰拾壹元整", convert(111))
	ass.Equal("壹佰零壹元整", convert(101))
	ass.Equal("壹仟壹佰拾壹元整", convert(1111))
	ass.Equal("壹仟壹佰零壹元整", convert(1101))
	ass.Equal("壹仟零拾壹元整", convert(1011))
	ass.Equal("壹万壹仟壹佰拾壹元整", convert(11111))
	ass.Equal("拾壹万壹仟壹佰拾壹元整", convert(111111))
	ass.Equal("贰拾壹万壹仟壹佰拾壹元整", convert(211111))
	ass.Equal("壹佰拾壹万壹仟壹佰拾壹元整", convert(1111111))
	ass.Equal("壹佰零壹万壹仟壹佰拾壹元整", convert(1011111))
	ass.Equal("壹仟壹佰零壹万壹仟壹佰拾壹元整", convert(11011111))
	ass.Equal("壹仟零拾壹万壹仟壹佰拾壹元整", convert(10111111))
	ass.Equal("壹亿壹仟零拾壹万壹仟壹佰拾壹元整", convert(110111111))
	ass.Equal("拾壹亿壹仟零拾壹万壹仟壹佰拾壹元整", convert(1110111111))
	ass.Equal("拾亿壹仟零拾壹万壹仟壹佰拾壹元整", convert(1010111111))
	ass.Equal("拾亿零拾壹万壹仟壹佰拾壹元整", convert(1000111111))
	ass.Equal("陆仟零柒元壹角肆分", convert(6007.14))
	ass.Equal("拾伍万壹仟壹佰贰拾壹元壹角伍分", convert(151121.15))
	ass.Equal("壹仟零拾元整", convert(1010.00))
}

func convert(f float64) (words string) {

	nf, ff := math.Modf(f)
	n := int(nf)
	frac := int(math.Round(ff * 100))

	yi := n / 1_0000_0000
	n %= 1_0000_0000
	wan := n / 1_0000
	n %= 1_0000

	if yi > 0 {
		words += xxxxToWords(yi) + "亿"
	}
	if wan > 0 {
		if yi > 0 && wan < 1000 {
			words += "零"
		}
		words += xxxxToWords(wan) + "万"
	}
	if yi > 0 && wan == 0 {
		words += "零"
	}

	words += xxxxToWords(n)
	if len(words) > 0 {
		words += "元"
	}
	if frac > 0 {
		words += fracToWords(frac)
	} else {
		words += "整"
	}
	return
	// 佰、仟、万、亿、元、角、分、零
}

func xxxxToWords(n int) (words string) {

	qian := n / 1000
	n %= 1000
	bai := n / 100
	n %= 100
	shi := n / 10
	n %= 10

	if qian > 0 {
		words += numToWord[qian] + "仟"
	}
	if bai > 0 {
		words += numToWord[bai] + "佰"
	}
	if shi > 0 {
		if qian >= 1 && bai == 0 {
			words += "零"
		}
		if shi == 1 {
			words += "拾"
		} else {
			words += numToWord[shi] + "拾"
		}
	}
	if n > 0 {
		if (qian > 0 || bai > 0) && shi == 0 {
			words += "零"
		}
		words += numToWord[n]
	}
	return
}

func fracToWords(n int) (words string) {
	if n >= 10 {
		jiao := n / 10
		n %= 10
		words += numToWord[jiao] + "角"
	}
	if n >= 1 {
		words += numToWord[n] + "分"
	}
	return
}
