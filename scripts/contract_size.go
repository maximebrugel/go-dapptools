package main

import (
	"os"
	"strconv"
	"dapptools-template/scripts/utils"
	"github.com/fatih/color"
)

func main() {
	if len(os.Args) != 3 {
		color.Red("Wrong args... need contract path and name")
		os.Exit(1)
	}
	size := utils.ContractSize(os.Args[1], os.Args[2])
	color.Cyan("%s bytes", strconv.Itoa(size))

	sizeLimit := 24576
	if sizeLimit > size {	
		color.Yellow("%s bytes left to reach the smart contract size limit of 24576 bytes.", strconv.Itoa(sizeLimit - size))
	} else {
		color.Yellow("%s bytes more than the smart contract size limit of 24576 bytes.", strconv.Itoa(size - sizeLimit))
	}
}
