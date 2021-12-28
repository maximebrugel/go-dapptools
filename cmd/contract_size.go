package cmd

import (
	"strconv"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"dapptools-template/cmd/utils"
)

var contractSizeCmd = &cobra.Command{
	Use:   "contract-size",
	Short: "Get the contract size of a smart contract",
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		printContractSize(args)
	},
}

func init() {
	rootCmd.AddCommand(contractSizeCmd)
}

// Args[0] : The smart contract path (for example : src/Greeter.sol)
// Args[1] : The contract name (for example : Greeter)
func printContractSize(args []string) {
	size := utils.ContractSize(args[0], args[1])
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