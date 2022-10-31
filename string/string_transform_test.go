package string

import (
	"fmt"
	"strings"
	"testing"
)

/*
nowcoder.com HJ29
描述
对输入的字符串进行加解密，并输出。

加密方法为：

当内容是英文字母时则用该英文字母的后一个字母替换，同时字母变换大小写,如字母a时则替换为B；字母Z时则替换为a；

当内容是数字时则把该数字加1，如0替换1，1替换2，9替换0；

其他字符不做变化。

解密方法为加密的逆过程。
数据范围：输入的两个字符串长度满足 1 \le n \le 1000 \1≤n≤1000  ，保证输入的字符串都是只由大小写字母或者数字组成
输入描述：
第一行输入一串要加密的密码
第二行输入一串加过密的密码

输出描述：
第一行输出加密后的字符
第二行输出解密后的字符

input:
abcdefg
BCDEFGH

output:
BCDEFGH
abcdefg
*/
func TestName(t *testing.T) {

	var sb strings.Builder

	plainText := "012789 ABCXYZ_abcxyz"
	for _, r := range plainText {
		sb.WriteRune(encrypt(r))
	}
	cipher := sb.String()
	fmt.Println(cipher)

	sb.Reset()

	for _, r := range cipher {
		sb.WriteRune(decrypt(r))
	}
	fmt.Println(sb.String())
}

func encrypt(r rune) rune {

	switch {
	case '0' <= r && r <= '9':
		return (r-'0'+1)%10 + '0'
	case 'a' <= r && r <= 'z':
		return (r-'a'+1)%26 + 'A'
	case 'A' <= r && r <= 'Z':
		return (r-'A'+1)%26 + 'a'
	default:
		return r
	}
}

func decrypt(r rune) rune {

	switch {
	case '0' <= r && r <= '9':
		return (r-'0'-1+10)%10 + '0'
	case 'a' <= r && r <= 'z':
		return (r-'a'-1+26)%26 + 'A'
	case 'A' <= r && r <= 'Z':
		return (r-'A'-1+26)%26 + 'a'
	default:
		return r
	}
}
