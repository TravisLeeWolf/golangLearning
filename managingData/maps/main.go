package main

import "fmt"

// More like a database than a dictionary
// type Product struct {
// 	id    int
// 	name  string
// 	price float64
// }
// Structs are good for clearly defined data where maps can be more dynamic

func main() {
	// Maps are basically like dictionaries in Python
	websites := map[string]string{
		"Apple":    "https://apple.com",
		"Google":   "https://google.com",
		"Facebook": "https://facebook.com",
		"Amazon":   "https://amazon.com",
		"X":        "https://x.com",
	}

	fmt.Println(websites["Apple"])
	websites["LinkedIn"] = "https://linkedin.com"
	delete(websites, "Facebook")

	fmt.Println(websites)

	siteRatings := make(map[string]float64, 5)
	siteRatings["Apple"] = 3.8
	siteRatings["Google"] = 4.2
	siteRatings["Facebook"] = 2.5

	for site, rating := range siteRatings {
		fmt.Println(site, rating)
	}

}
