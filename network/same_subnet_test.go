package network

import (
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestCheckSameSubnet(t *testing.T) {

	ass := assert.New(t)
	ass.EqualValues(1, checkSameSubnet("255.255.255.0", "192.168.224.256", "192.168.10.4"))
}

func checkSameSubnet(dotMask, dotIp1, dotIp2 string) int {
	mask := dotMaskToInt(dotMask)
	ip1 := dotIpToInt(dotIp1)
	ip2 := dotIpToInt(dotIp2)
	if mask < 0 || ip1 < 0 || ip2 < 0 {
		return 1
	} else if ip1&mask == ip2&mask {
		return 0
	} else {
		return 2
	}
}

func dotIpToInt(dotIp string) int64 {
	ipSplits := strings.Split(dotIp, ".")
	var ipInt int64 = 0
	if len(ipSplits) == 4 {
		for i := 0; i < 4; i++ {
			ipSplitNum, err := strconv.Atoi(ipSplits[i])
			if err == nil {
				if ipSplitNum < 0 || ipSplitNum > 255 {
					return -1
				}
			} else {
				return -1
			}
			ipInt += int64(ipSplitNum) << (8 * (3 - i))
		}
		return ipInt
	} else {
		return -1
	}
}

func dotMaskToInt(dotMask string) int64 {
	maskInt := dotIpToInt(dotMask)
	if maskInt == 0 ||
		maskInt == math.MaxUint32 ||
		(maskInt|(maskInt-1) != math.MaxUint32) {
		return -1
	} else {
		return maskInt
	}
}
