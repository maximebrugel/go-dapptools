package main

import (
	"dapptools-template/scripts/utils"
	"os"
	"strconv"
	"github.com/fatih/color"
)

// Command to get the contract size of a smart contract
// Takes two arguments :
// PATH : The smart contract path (for example : src/Greeter.sol)
// NAME : The contract name (for example : Greeter)
func main() {
	if len(os.Args) != 3 {
		color.Red("Wrong args... need contract path and name")
		os.Exit(1)
	}
	size := utils.ContractSize(os.Args[1], os.Args[2])
	color.Cyan("%s bytes", strconv.Itoa(size))

	sizeLimit := 24576
	if sizeLimit > size {
		color.Yellow("%s bytes left to reach the smart contract size limit of 24576 bytes.",
			strconv.Itoa(sizeLimit-size))
	} else {
		color.Yellow("%s bytes more than the smart contract size limit of 24576 bytes.", 
			strconv.Itoa(size-sizeLimit))
	}
}
