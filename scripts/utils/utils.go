package utils

import (
	"github.com/buger/jsonparser"
	"os"
	"io/ioutil"
	"fmt"
)

// Return the contract size of a given contract
func ContractSize(path string, name string) int {
	data := readDappOutput()

	strContent, err := jsonparser.GetString(data, "contracts", path, name, "evm", "bytecode", "object")

	if err != nil {
		os.Exit(1)
	}

	return len([]rune(strContent)) / 2
}


// DappTools creates a json file with `dapp build` 
// It returns the content ([]byte) of this json file
func readDappOutput() []byte {
	data, err := ioutil.ReadFile("./out/dapp.sol.json")

	if err != nil {
		fmt.Println(err);
		fmt.Println("Please compile using dapp build")
		os.Exit(1)
	}

	return data
}