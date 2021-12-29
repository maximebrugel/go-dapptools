package cmd

import (
	"os"
	"os/exec"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/joho/godotenv"
	"dapptools-template/cmd/utils"
)

var deployEstimateCmd = &cobra.Command{
	Use:   "deploy-estimate",
	Short: "Get the deployment gas cost of a smart contract",
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		deployEstimate(args)
	},
}

func init() {
	rootCmd.AddCommand(deployEstimateCmd)
}

func deployEstimate(args []string) {
	err := godotenv.Load()
  	if err != nil {
		color.Red("Error loading .env file")
		os.Exit(1)
  	}

  	rpcUrl := os.Getenv("ETH_RPC_URL")
	sethCmd := os.Getenv("HOME") + "/.nix-profile/bin/seth"

	data := utils.ReadDappOutput()
	abi := utils.GetAbi(data, args[0], args[1])
 
	sig, err := exec.Command(sethCmd, "--abi-constructor", abi).Output()
	if err != nil {
		sig = []byte("constructor()")
	}

	bytecode := utils.GetBytecode(data, args[0], args[1])
	deployArgs := args[2:]

	params := append([]string{"estimate", "--create", bytecode, string(sig)}, deployArgs...)
	params = append(params, "--rpc-url", rpcUrl)

	gas, err := exec.Command(sethCmd, params...).Output()
	if err != nil {
		color.Red(string(gas))
		os.Exit(1)
	}

	color.Cyan("Gas to deploy : %s", gas)
}