package network

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var aMin = 1 << 24
var aMax = 127<<24 - 1
var bMin = 128 << 24
var bMax = 192<<24 - 1
var cMin = 192 << 24
var cMax = 224<<24 - 1
var dMin = 224 << 24
var dMax = 240<<24 - 1
var eMin = 240 << 24
var eMax = 256<<24 - 1

var privateIpMin1 = 10 << 24
var privateIpMax1 = 11<<24 - 1
var privateIpMin2 = 172<<24 + 16<<16
var privateIpMax2 = 172<<24 + 32<<16 - 1
var privateIpMin3 = 192<<24 + 168<<16
var privateIpMax3 = 192<<24 + 169<<16 - 1

var ignoreIpMin1 = 0
var ignoreIpMax1 = 1<<24 - 1
var ignoreIpMin2 = 127 << 24
var ignoreIpMax2 = 128<<24 - 1

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	aCount, bCount, cCount, dCount, eCount := 0, 0, 0, 0, 0
	illegalIpCount := 0
	privateIpCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, "~")
		ip := splits[0]
		mask := splits[1]
		var ipInt int = parseIp(ip)
		var maskInt int = parseMask(mask)
		if ipInt < 0 {
			illegalIpCount++
		} else if (ignoreIpMin1 <= ipInt && ipInt <= ignoreIpMax1) || (ignoreIpMin2 <= ipInt && ipInt <= ignoreIpMax2) {
			// ignore
		} else if maskInt < 0 {
			illegalIpCount++
		} else {
			var maskIntXor int = maskInt ^ math.MaxUint32
			var ipIntRangeMin int = int(ipInt)
			var ipIntRangeMax int = int(ipInt) | maskIntXor
			switch {
			case aMin <= ipIntRangeMin && ipIntRangeMin <= aMax &&
				aMin <= ipIntRangeMax && ipIntRangeMax <= aMax:
				aCount++
			case bMin <= ipIntRangeMin && ipIntRangeMin <= bMax &&
				bMin <= ipIntRangeMax && ipIntRangeMax <= bMax:
				bCount++
			case cMin <= ipIntRangeMin && ipIntRangeMin <= cMax &&
				cMin <= ipIntRangeMax && ipIntRangeMax <= cMax:
				cCount++
			case dMin <= ipIntRangeMin && ipIntRangeMin <= dMax &&
				dMin <= ipIntRangeMax && ipIntRangeMax <= dMax:
				dCount++
			case eMin <= ipIntRangeMin && ipIntRangeMin <= eMax &&
				eMin <= ipIntRangeMax && ipIntRangeMax <= eMax:
				eCount++
			}
			switch {
			case privateIpMin1 <= ipIntRangeMin && ipIntRangeMin <= privateIpMax1 &&
				privateIpMin1 <= ipIntRangeMax && ipIntRangeMax <= privateIpMax1:
				privateIpCount++
			case privateIpMin2 <= ipIntRangeMin && ipIntRangeMin <= privateIpMax2 &&
				privateIpMin2 <= ipIntRangeMax && ipIntRangeMax <= privateIpMax2:
				privateIpCount++
			case privateIpMin3 <= ipIntRangeMin && ipIntRangeMin <= privateIpMax3 &&
				privateIpMin3 <= ipIntRangeMax && ipIntRangeMax <= privateIpMax3:
				privateIpCount++
			}
		}
	}
	fmt.Printf("%d %d %d %d %d %d %d", aCount, bCount, cCount, dCount, eCount, illegalIpCount, privateIpCount)

}

func parseIp(ip string) int {
	ipSplits := strings.Split(ip, ".")
	var ipInt int = 0
	if len(ipSplits) == 4 {
		for i := 3; i >= 0; i-- {
			ipSplitNum, err := strconv.Atoi(ipSplits[i])
			if err == nil {
				if ipSplitNum < 0 || ipSplitNum > 255 {
					return -1
				}
			} else {
				return -1
			}
			ipInt += int(ipSplitNum) << (8 * (3 - i))
		}
	}
	return ipInt
}

func parseMask(mask string) int {
	maskSplits := strings.Split(mask, ".")
	var maskInt int = 0
	if len(maskSplits) == 4 {
		maskBinaryString := ""
		for i := 3; i >= 0; i-- {
			maskPart := maskSplits[i]
			maskPartNum, err := strconv.Atoi(maskPart)
			if err == nil && 0 <= maskPartNum && maskPartNum <= 255 {
				maskBinaryString = fmt.Sprintf("%08b", maskPartNum) + maskBinaryString
			} else {
				return -1
			}
			maskInt += int(maskPartNum) << (8 * (3 - i))
		}
		// now the ip address is legal

		var found0, found1 bool
		for _, bin := range maskBinaryString {
			if bin == '0' {
				found0 = true
				if !found1 { // 0 before 1
					return -1
				}
			} else { // 1
				found1 = true
				if found0 { // 0 before 1
					return -1
				}
			}
		}
		if !found0 { // no 0
			return -1
		}
		// now the mask is legal
	}
	return maskInt
}
