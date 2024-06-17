package main

import "fmt"

func main() {
	prices := []float64{5.99}
	prices = append(prices, 9.99)
	prices = prices[1:]

	otherPrices := []float64{1.99, 2.99, 3.99}
	prices = append(prices, otherPrices...)

	fmt.Println(prices)

	// Using make
	names := make([]string, 2, 5)
	names[0] = "Bob"

	names = append(names, "John")

	for index, name := range names {
		fmt.Println(index, name)
	}
}

// func basic() {
// 	productNames := [5]string{"Apple", "Orange", "Banana", "Pear", "Grape"}
// 	prices := [5]float64{9.99, 19.99, 29.99, 39.99, 49.99}

// 	fmt.Println(prices)
// 	fmt.Println(productNames)

// 	// Slices
// 	featuredProducts := productNames[1:4]
// 	fmt.Println(featuredProducts)

// 	highPrices := prices[3:]
// 	fmt.Println(highPrices)
// 	// NOTE: Can't use negative indexes in slices

// 	// NOTE: Modifying the slice will modify the original array
// 	highPrices[0] = 99.99
// 	fmt.Println(highPrices)
// 	fmt.Println(prices)

// 	fmt.Println(len(featuredProducts), cap(featuredProducts))
// 	fmt.Println(len(highPrices), cap(highPrices))
// 	// More elements available to the right
// }
