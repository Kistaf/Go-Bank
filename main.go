package main

import (
	"fmt"

	banks "github.com/Kistaf/banks"
	types "github.com/Kistaf/types"
)

func main() {
	bankAccounts := []types.Banker{
		banks.NewStandardBank(500),
		banks.NewBitcoinBank(50),
	}

	for _, b := range bankAccounts {
		fmt.Println(b.Balance())
		err := b.Withdraw(35)
		if err != nil {
            fmt.Println(err)
            continue
		}
		fmt.Println(b.Balance())
	}
}
