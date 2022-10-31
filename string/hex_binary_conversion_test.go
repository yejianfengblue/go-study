package string

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

// nowcoder.com HJ30
func TestHexBinConversion(t *testing.T) {

	hex := 'a'
	if ('0' <= hex && hex <= '9') || ('a' <= hex && hex <= 'f') || ('A' <= hex && hex <= 'F') {
		decimalValue, _ := strconv.ParseUint(string(hex), 16, 0)
		binaryString := fmt.Sprintf("%04b", decimalValue)
		reversedBinaryRunes := []rune(binaryString)
		sort.Slice(reversedBinaryRunes, func(i, j int) bool {
			return i > j
		})
		reversedDecimalValue, _ := strconv.ParseUint(string(reversedBinaryRunes), 2, 0)
		reversedHex := fmt.Sprintf("%X", reversedDecimalValue)
		fmt.Printf("hex %q -> dec %d -> bin %s -> rev bin %s -> rev dec %d -> rev hex %s \n", hex,
			decimalValue, binaryString, string(reversedBinaryRunes), reversedDecimalValue, reversedHex)
	}
}
