# Interchain Security

[![Go Report Card](https://goreportcard.com/badge/github.com/cosmos/interchain-security)](https://goreportcard.com/report/github.com/cosmos/interchain-security)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=cosmos_interchain-security&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=cosmos_interchain-security)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=cosmos_interchain-security&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=cosmos_interchain-security)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=cosmos_interchain-security&metric=bugs)](https://sonarcloud.io/summary/new_code?id=cosmos_interchain-security)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=cosmos_interchain-security&metric=ncloc)](https://sonarcloud.io/summary/new_code?id=cosmos_interchain-security)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=cosmos_interchain-security&metric=coverage)](https://sonarcloud.io/summary/new_code?id=cosmos_interchain-security)

**interchain-security** contains a working and in-production implementation of the Interchain Security (ICS) protocol. ICS is an open sourced IBC application which allows cosmos blockchains to lease their proof-of-stake security to one another.

For more details on the **Interchain Security protocol**, take a look at the [docs](https://cosmos.github.io/interchain-security/) or [technical specification](https://github.com/cosmos/ibc/blob/main/spec/app/ics-028-cross-chain-validation/README.md).

For a list of **currently active releases**, see [RELEASES.md](./RELEASES.md#version-matrix).

For a list of **major ICS features** available in the currently active releases, see [FEATURES.md](./FEATURES.md).

## Instructions

**Prerequisites**

```bash
## For OSX or Linux

# go 1.21 (https://formulae.brew.sh/formula/go)
brew install go@1.21
# jq (optional, for testnet) (https://formulae.brew.sh/formula/jq)
brew install jq
# docker (optional, for integration tests, testnet) (https://docs.docker.com/get-docker/)

```

**Installing and running binaries**

```bash
# install interchain-security-pd and interchain-security-cd binaries
make install
# run provider
interchain-security-pd
# run consumer
interchain-security-cd
# (if the above fail, ensure ~/go/bin on $PATH)
export PATH=$PATH:$(go env GOPATH)/bin
```

Inspect the [Makefile](./Makefile) if curious.

## Testing

See [testing docs](./TESTING.md).

## Learn more

- [IBC Docs](https://ibc.cosmos.network/)
- [IBC Protocol](https://ibcprotocol.org/)
- [IBC Specs](https://github.com/cosmos/ibc)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/cosmosnetwork)

## My Instructions ## (if the main instruction from above doesnt work)

# in Makefile replace:

BUILD_TARGETS := build

build: BUILD_ARGS=-o $(BUILDDIR)/

$(BUILD_TARGETS): go.sum $(BUILDDIR)/
go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./...

# with: (optionally with interchain-security-cdd, interchain-security-sd if needed)

.PHONY: build
build: $(BUILDDIR)/
go build -o $(BUILDDIR)/interchain-security-pd ./cmd/interchain-security-pd
go build -o $(BUILDDIR)/interchain-security-cd ./cmd/interchain-security-cd
go build -o $(BUILDDIR)/interchain-security-cdd ./cmd/interchain-security-cdd
go build -o $(BUILDDIR)/interchain-security-sd ./cmd/interchain-security-sd

# next run: make build

# right issue: check ->

ls -ld /Users/markokuncic/Desktop/Cosmos_SDK/interchain-security/build

# if rights not given run:

sudo chown -R markokuncic:staff /Users/markokuncic/Desktop/Cosmos_SDK/interchain-security/build

# this should give access rights to the build folder

---

**_Notes_**

# Denominator

- 1 MAANY = 1.000.000 stake (smallest unit)

**_ Instruction _**

1. Initialize as described above -> should have build folder with provider chain
2. Initialize the genesis.json:
   - ./binary/interchain-security-pd init <node-name> --chain-id <chain-ID>
3. Initialize gentx file and validator account with self delegation
   - ./init-gentx.sh
4. Populate accounts:
   - ./prepare-genesis.sh
5. Reset Inflation
   - set all values of "inflation" to "0.00000000000000000"
6. Activate gas fees
   - set minimum-gas-prices = "1stake" in app.toml
7. Activate block rewards
   - How?
8. Set ticker info
   - ?
9. Create burn wallet:
   - ?
10. Burning logic:
    - adjust params.community_tax in the gensis.json file
    - those go to the community-pool wallet

**_Other Commands_**

# To run blockchain from block 0:

- delete all files from data folder in root
- reset priv_validator_state.json back to 0, 0, 0 without "signature" and "signbytes"

# To reset accounts and genesis file:

- run script: ./reset-genesis.sh

# Run blockchain

- ./build/interchain-security-pd start

**_Token Distribution_**

# Populate tokendistribution

- reserve tokens for validator (e.g. 5%) -> register like above, i.e. with gentx command for an already created validator address

- Unlocked Wallets

  - ALM (10%)
  - Liquidity (5%)
  - Validators (5%)
  - Public (6%)

- Linear Vesting Schedule Wallets

  - Foundation (10%)
  - Community and Dev (7%)
  - Team (13%)
  - Advisors (5%)
  - Investor Wallets (see below) (9%)

- Investor Wallets

  - Pre-Seed/Seed/Private (9%)
    -> need a wallet logic, such that investors can create their own addresses (this even works when the )

    -> See: https://chatgpt.com/c/6737ce92-dfc8-800c-8618-c5f3f4d6034b
    -> Keyword: How Wallets Work Under the Hood

- Minting Logic
  - Minting Wallet (30%)
