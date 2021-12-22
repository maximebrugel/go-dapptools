# include .env file and export its env vars
args = $(filter-out $@,$(MAKECMDGOALS))

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
estimate:; go run ./scripts/contract_size.go $(call args)