package utils

import (
	"github.com/buger/jsonparser"
	"os"
	"io/ioutil"
	"github.com/fatih/color"
)

// Return the contract size of a given contract
func ContractSize(path string, name string) int {
	data := ReadDappOutput()
	bytecode := GetBytecode(data, path, name)
	return len([]rune(bytecode)) / 2
}

func GetBytecode(data []byte, path string, name string) string {
	bytecode, err := jsonparser.GetString(data, "contracts", path, name, "evm", "bytecode", "object")
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
	return bytecode
}

func GetAbi(data []byte, path string, name string) string {
	abi, _, _, err := jsonparser.Get(data, "contracts", path, name, "abi")
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
	return string(abi)
}


// DappTools creates a json file with `dapp build` 
// It returns the content ([]byte) of this json file
func ReadDappOutput() []byte {
	data, err := ioutil.ReadFile("./out/dapp.sol.json")

	if err != nil {
		color.Red(err.Error());
		color.Red("Please compile using dapp build")
		os.Exit(1)
	}

	return data
}