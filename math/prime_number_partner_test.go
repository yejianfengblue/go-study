package math

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

var nums []int
var indexes []int

// n * n 2D array, value 0 means no cache, 1 means is prime, -1 mean not prime
var cache [][]int

func TestIsPrimeNumberPartner(t *testing.T) {

	ass := assert.New(t)

	nums = []int{2, 5, 6, 13}
	setup(nums)
	ass.EqualValues(2, CountPrimeNumberPartner(indexes))

	nums = []int{3, 6}
	setup(nums)
	ass.EqualValues(0, CountPrimeNumberPartner(indexes))

	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	for i, s := range strings.Split(line, " ") {
		nums[i], _ = strconv.Atoi(s)
	}
}

func setup(nums []int) {
	cache = make([][]int, len(nums))
	for i := 0; i < len(cache); i++ {
		cache[i] = make([]int, len(nums))
	}

	indexes = make([]int, len(nums))
	for i := range indexes {
		indexes[i] = i
	}
}

func CountPrimeNumberPartner(indexes []int) int {
	if len(indexes) == 2 {
		if getPrimeCache(indexes[0], indexes[1]) {
			return 1
		} else {
			return 0
		}
	}
	countMax := 0
	// now at least 4 numbers
	// pick 2 numbers first, recursively call with the remaining numbers
	for i := 0; i < len(indexes)-1; i++ {
		for j := i + 1; j < len(indexes); j++ {
			// is ij prime
			ijCount := 0
			if getPrimeCache(indexes[i], indexes[j]) {
				ijCount++
			}
			// count the prime number of the remaining
			remainingIndexes := make([]int, 0, len(indexes)-2)
			for remainingI := i + 1; remainingI < j; remainingI++ {
				remainingIndexes = append(remainingIndexes, indexes[remainingI])
			}
			if j < len(indexes)-1 {
				remainingIndexes = append(remainingIndexes, indexes[j+1:]...)
			}
			thisCount := ijCount + CountPrimeNumberPartner(remainingIndexes)
			if thisCount > countMax {
				countMax = thisCount
				fmt.Printf("countMax %d indexi=%d indexj=%d remainingIndexes=%v\n",
					countMax, indexes[i], indexes[j], remainingIndexes)
			}
		}
	}
	return countMax
}

func getPrimeCache(i, j int) bool {
	if cache[i][j] == 0 { // if no cache, calc it
		if IsPrime(nums[i] + nums[j]) {
			cache[i][j] = 1
		} else {
			cache[i][j] = -1
		}
		return cache[i][j] > 0
	} else {
		return cache[i][j] > 0
	}
}

func IsPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 5; i < limit+1; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
