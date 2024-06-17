package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		fmt.Println("---#---#---")
		panic("Fatal error, app died...")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7 at: ", randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Please deposit an amount greater than 0. Please try again")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Your new balance is:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)
			if (accountBalance - withdrawalAmount) < 0 {
				fmt.Println("You do not have enough funds!")
			} else {
				accountBalance -= withdrawalAmount
				fmt.Println("Your new balance is:", accountBalance)
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			}
		default:
			fmt.Println("---#---#---")
			fmt.Println("Thank you for using Go Bank")
			fmt.Println("Goodbye!")
			return
		}
	}
}
