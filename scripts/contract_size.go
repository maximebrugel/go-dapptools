package main

import (
    "fmt"
	"os"
	"dapptools-template/scripts/utils"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Wrong args... need contract path and name")
		os.Exit(1)
	}
	size := utils.ContractSize(os.Args[1], os.Args[2])
	fmt.Println(size, " bytes")
	fmt.Println(24576 - size, " bytes left to reach the smart contract size limit of 24576 bytes.")
}
