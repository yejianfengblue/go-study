package main

import (
	"fmt"
	"strings"
)

/*
max sum of value * weight
*/
func main() {

	money, itemAmount := 0, 0
	for {
		scanN, _ := fmt.Scan(&money, &itemAmount)
		if scanN == 0 {
			break
		} else {

			fmt.Printf("%d\n", UseStruct(money, itemAmount)*10)
			fmt.Printf("%d\n", UseStructSmall(money, itemAmount)*10)

		}
	}
}

func UseStruct(money int, itemAmount int) int {

	money /= 10 // because all item prices are multiple of 10, divide by 10 to reduce the array size
	// when make the array, all the struct members are initialized with zero-values
	items := make([]Item, itemAmount+1)
	for i := 1; i <= itemAmount; i++ {
		v, p, q := 0, 0, 0
		_, _ = fmt.Scan(&v, &p, &q)
		v /= 10 // divide by 10, see the money comment
		if q == 0 {
			// this item is main, put it into the items
			items[i].Price = v
			items[i].Weight = v * p
			items[i].Main = q
		} else {
			// this item is accessory, put it into the accessories array of main
			main := &items[q]
			if main.Accessories == nil {
				main.Accessories = make([]Item, itemAmount)[:0]
			}
			main.Accessories = append(main.Accessories, Item{v, v * p, q, nil})
			// also mark items[i] to be accessory, so we skip processing it later
			items[i].Main = q
		}
	}
	//for i, item := range items {
	//	fmt.Printf("%d %v\n", i, item)
	//}
	// tabulation[i][j] means given j capacity (money), when buy the first i items, the best value (satisfaction)
	// see 0/1 Knapsack Problem Dynamic Programming https://www.youtube.com/watch?v=8LusJS5-AGo
	tabulation := make([][]int, itemAmount+1)
	for i := 0; i <= itemAmount; i++ {
		tabulation[i] = make([]int, money+1)
	}
	// when buy 0 items, the satisfaction is 0
	for j := 0; j <= money; j++ {
		tabulation[0][j] = 0
	}

	tabi := 0
	for i := 1; i <= itemAmount; i++ {
		item := items[i]
		if item.Main == 0 {
			tabi++
			main := item
			for j := 0; j <= money; j++ {
				// if j (money) < price, the best value is same as above row
				if j < main.Price {
					tabulation[tabi][j] = tabulation[tabi-1][j]
				} else {
					// given money 10, if buy this main whose price is 7, remain 3
					// ask what is the best value of money 3 during the previous processing
					// compare their sum with if not buy this main
					ifBuy := main.Weight + tabulation[tabi-1][j-item.Price]
					ifNotBuy := tabulation[tabi-1][j]
					if ifBuy > ifNotBuy {
						tabulation[tabi][j] = ifBuy
					} else {
						tabulation[tabi][j] = ifNotBuy
					}
				}
				//fmt.Printf("\n\n%s", Print2DIntArray(tabulation))
			}
			accessories := main.Accessories
			for acci := 0; acci < len(accessories); acci++ {
				acc := accessories[acci]
				tabi++
				for j := 0; j <= money; j++ {
					if j < main.Price+acc.Price { // money < main + this accessory
						tabulation[tabi][j] = tabulation[tabi-1-acci][j]
					} else { // money >= main + this accessory
						// when acci = 0, tabulation[tabi-1-acci] is main
						// tabulation[tabi-1-acci-1] is not buy main
						ifBuyMainAndThisAccessory := main.Weight + acc.Weight + tabulation[tabi-1-acci-1][j-item.Price-acc.Price]
						ifNotBuy := tabulation[tabi-1][j]
						if ifBuyMainAndThisAccessory > ifNotBuy {
							tabulation[tabi][j] = ifBuyMainAndThisAccessory
						} else {
							tabulation[tabi][j] = ifNotBuy
						}
						if acci == 1 { // try to buy main and two accessories
							mainAndTwoAccessoriesPrice := main.Price +
								accessories[0].Price + accessories[1].Price
							if j >= mainAndTwoAccessoriesPrice {
								ifBuyMainAndTwoAccessories := main.Weight +
									accessories[0].Weight + accessories[1].Weight +
									tabulation[tabi-1-acci-1][j-mainAndTwoAccessoriesPrice]
								if ifBuyMainAndTwoAccessories > ifNotBuy {
									tabulation[tabi][j] = ifBuyMainAndTwoAccessories
								} else {
									tabulation[tabi][j] = ifNotBuy
								}
							}
						}
					}
					fmt.Printf("\n\n%s", Print2DIntArray(tabulation))
				}
			}
		} // if item.Main == 0
	}

	//fmt.Printf("\n\n%s", Print2DIntArray(tabulation))
	return tabulation[itemAmount][money]
}

func UseStructSmall(money int, itemAmount int) int {

	var tabulation [3200]int
	money /= 10 // because all item prices are multiple of 10, divide by 10 to reduce the array size
	// when make the array, all the struct members are initialized with zero-values
	items := make([]Item, itemAmount+1)
	for i := 1; i <= itemAmount; i++ {
		v, p, q := 0, 0, 0
		_, _ = fmt.Scan(&v, &p, &q)
		v /= 10 // divide by 10, see the money comment
		if q == 0 {
			// this item is main, put it into the items
			items[i].Price = v
			items[i].Weight = v * p
			items[i].Main = q
		} else {
			// this item is accessory, put it into the accessories array of main
			main := &items[q]
			if main.Accessories == nil {
				main.Accessories = make([]Item, 2)[:0]
			}
			main.Accessories = append(main.Accessories, Item{v, v * p, q, nil})
			// also mark items[i] to be accessory, so we skip processing it later
			items[i].Main = q
		}
	}
	for i, item := range items {
		fmt.Printf("%d %v\n", i, item)
	}

	for i := 1; i <= itemAmount; i++ {
		item := items[i]
		if item.Main == 0 {
			main := item
			mainPrice, mainWeight := main.Price, main.Weight
			mainAccessory0Price, mainAccessory0Weight := -1, -1   // buy main and accessory 0
			mainAccessory1Price, mainAccessory1Weight := -1, -1   // buy main and accessory 1
			mainAccessory01Price, mainAccessory01Weight := -1, -1 // buy main and accessory 0 and 1
			if len(main.Accessories) > 0 {
				mainAccessory0Price = main.Price + main.Accessories[0].Price
				mainAccessory0Weight = main.Weight + main.Accessories[0].Weight
				if len(main.Accessories) > 1 {
					mainAccessory1Price = main.Price + main.Accessories[1].Price
					mainAccessory1Weight = main.Weight + main.Accessories[1].Weight
					mainAccessory01Price = main.Price + main.Accessories[0].Price + main.Accessories[1].Price
					mainAccessory01Weight = main.Weight + main.Accessories[0].Weight + main.Accessories[1].Weight
				}
			}
			for j := money; j > 0; j-- {
				if j >= mainPrice {
					tabulation[j] = Max(mainWeight+tabulation[j-mainPrice], tabulation[j])
					if j >= mainAccessory0Price && mainAccessory0Weight > 0 {
						tabulation[j] = Max(mainAccessory0Weight+tabulation[j-mainAccessory0Price], tabulation[j])
					}
					if j >= mainAccessory1Price && mainAccessory1Weight > 0 {
						tabulation[j] = Max(mainAccessory1Weight+tabulation[j-mainAccessory1Price], tabulation[j])
					}
					if j >= mainAccessory01Price && mainAccessory01Weight > 0 {
						tabulation[j] = Max(mainAccessory01Weight+tabulation[j-mainAccessory01Price], tabulation[j])
					}
				}
				//fmt.Printf("\n\n%v", tabulation[0:money+1])
			}
		} // if item.Main == 0
	}

	return tabulation[money]
}

type Item struct {
	Price       int
	Weight      int
	Main        int
	Accessories []Item
}

func Print2DIntArray(array [][]int) string {

	var sb strings.Builder
	sb.WriteString("\n  ")
	for i := 0; i < len(array[0]); i++ {
		sb.WriteString(fmt.Sprintf("%6d", i))
	}
	sb.WriteString("\n")
	for i, v := range array {
		sb.WriteString(fmt.Sprintf("%2d%v \n", i, IntArrayToString(v)))
	}

	return sb.String()
}

func IntArrayToString(array []int) string {
	var sb strings.Builder
	for i := 0; i < len(array); i++ {
		sb.WriteString(fmt.Sprintf("%6d", array[i]))
	}
	return sb.String()
}

func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
