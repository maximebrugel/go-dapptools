# (-include to ignore error if it does not exist)
-include .env

install: solc update npm

# dapp deps
update:; dapp update

# npm deps for linting etc.
npm:; yarn install

# install solc version
# example to install other versions: `make solc 0_8_2`
SOLC_VERSION := 0_8_10
solc:; nix-env -f https://github.com/dapphub/dapptools/archive/master.tar.gz -iA solc-static-versions.solc_${SOLC_VERSION}

# Build & test
build  :; dapp build
test   :; dapp test # --ffi # enable if you need the `ffi` cheat code on HEVM
snapshot :; dapp snapshot
coverage :; dapp test --coverage
clean  :; dapp clean
lint   :; yarn run lint
prettier :; yarn prettier

# Scripts

# Estimate the contract size
# PATH : The path to the contract (for example : "src/Greeter.sol")
# NAME : The contract name (for example : "Greeter")
contract-size:; go run main.go contract-size $(PATH) $(NAME)