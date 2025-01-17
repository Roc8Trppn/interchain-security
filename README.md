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

4. Run script

- go run get_blockrewards_address.go
- copy address
- add module account into genesis.json
- update balance and total supply

4. Change Prefix "cosmos"

5. Populate accounts:
   - ./prepare-genesis.sh
6. Reset Inflation
   - set all values of "inflation" to "0.00000000000000000"
7. Activate gas fees
   - set minimum-gas-prices = "1stake" in app.toml
8. Activate block rewards
   - How?
9. Set ticker info
   - ?
10. Create burn wallet:
    - ?
11. Burning logic:
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

**_Qeuery_**

{
"@type": "/cosmos.auth.v1beta1.ModuleAccount",
"base_account": {
"address": "cosmos1kdsm4jzhnrck2ucykhrj8lhhayp3am3s6y6uzp",
"pub_key": null,
"account_number": "0",
"sequence": "0"
},
"name": "blockrewards",
"permissions": ["minter"]
},
./build/interchain-security-pd query bank balances cosmos1kdsm4jzhnrck2ucykhrj8lhhayp3am3s6y6uzp

**_CCV_**

- Consumer:
  -> ccvconsumer.params.enabled = true
  -> unbonding_period the same as in Provider
- Provider
  -> (PRE LAUNCH)

  -

- Mnemonic for Hermes Account
  ocean fit off offer amused pond tissue mixed motion pumpkin jacket road balcony brisk expect battle away destroy reject can add behind eye promote
- address: cosmos1lu5xltueklecryv3878qjgfjfjv5ezre3cf37t
  name: hermes-key
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AgjPhErSK2H1lLhF2UwQwQhua23EeoF7P8qxsYUg451L"}'
  type: local

**_Trasnactions_**

**_Governance_**

- JSON has to follow the structure as in draft_proposal

./build/interchain-security-pd tx gov submit-proposal ./test-proposal.json \
--from cosmos1yulh6xh93ld76rzdleek5vvmsh4fw4sdaempra \
--chain-id provider-chain \
--gas auto \
--fees 500stake

**_Check Transaction Hsash_**
./build/interchain-security-pd query tx 68EA148480D97C7BB7DF977BB69951CB736E3FB57A06641AC38D59F0D11304D9

**Steps to create new Consumer chain proposal**

1. Create a new consumer chain via ./build/interchain-security-cd init
1. Change ports in toml files:

- [address = "localhost:9091"]
- [pprof_laddr = "localhost:6062"]
- [node = "tcp://localhost:26657"]
- [p2p] # Address to listen for incoming connections
  laddr = "tcp://0.0.0.0:26659"
- [rpc] laddr = "tcp://127.0.0.1:26658"

2. Prepare json for provider settings of consumer chain (See ./proposals/create_consumer.json for example of flag params)
3. Run: ./build/interchain-security-pd tx provider create-consumer </path/o/json> --from=<address> --gas 200000 --fees 5000stake --chain-id provider-chain -> **_this should be a valid transaction with hash_**
4. Check out Hash: ./build/interchain-security-pd query tx <tx-hash> -> raw_logs should be empty Check value of key "consumer_phase" -> should see
5. Query provider: ./build/interchain-security-pd query provider list-consumer-chains -> "phase" should be ["CONSUMER_CHAIN_REGISTERED"]

6. Validator Opt-In: at least one validator must opt-in before [spawn_time] ends
   -> ./build/interchain-security-pd tx provider opt-in
   -> find public key via: ./build/interchain-security-pd query staking validators

7.

8. The gensis.json of the Consumer Chain must be valid to proceed
   https://github.com/hyphacoop/ics-testnets/blob/main/docs/Consumer-Chain-Start-Process.md

- Note: genesis file must look like in genesis_test.go
- ## Parameters to customize:

  [params.]
  -> [unbonding_period] in genesis.json has to match the one in the provider settings (see create_consumer.json) (TODO: check format in json, can be the same as in genesis.json?)
  -> [consumer_id] has to match the provider settings

  [provider.client_state.]
  -> [chain_id] must be the PROVIDER's chain id!
  -> [trusting_period] should be the same as in provider settings

  [provider.consensus_state.]
  -> [timestamp] Replace with the timestamp of the latest block on the provider chain (UTC format).
  -> [root.hash] Replace with the Merkle root hash of the latest block on the provider chain
  => ./build/interchain-security-pd status
  => get latest block height, timestamp, Merkle root

  [provider.initial_val_set.]
  -> [pub_key] Use the public keys of validators from the provider chain.
  -> [power] Define the voting power of each validator from the provider chain.
  => get from ./build/interchain-security-pd query staking validators

8. Next run: ./build/interchain-security-pd tx provider update-consumer <path/to/update\*comsumer.json> **here we provider the chain_id that was given after REGISTRATION**
   -> If transaction passes, the state should be [CONSUMER_CHAIN_INITIALIZED]

- NOTE: Chain can start at "spawn_time" (see output above) -> if chain is tried to be run before we should get
  -> [INFO] Consumer chain initialized. Waiting for spawn_time: <spawn_time>

9. Prepare Hermes

- check ~/.hermes/config.toml for settings
- in terminal in root, run [hermes start]

### Example Genesis

```
"ccvconsumer": {
      "params": {
        "enabled": true,
        "blocks_per_distribution_transmission": "1000",
        "distribution_transmission_channel": "",
        "provider_fee_pool_addr_str": "",
        "ccv_timeout_period": "2419200s",
        "transfer_timeout_period": "3600s",
        "consumer_redistribution_fraction": "0.75",
        "historical_entries": "10000",
        "unbonding_period": "1728000s",
        "reward_denoms": [],
        "provider_reward_denoms": [],
        "retry_delay_period": "3600s",
        "consumer_id": "0"
      },
      "provider": {
        "client_state": {
          "chain_id": "provider-chain",
          "trust_level": {
            "numerator": "1",
            "denominator": "3"
          },
          "trusting_period": "1197504s",
          "unbonding_period": "1814400s",
          "max_clock_drift": "10s",
          "frozen_height": null,
          "latest_height": {
            "revision_number": "0",
            "revision_height": "24"
          },
          "proof_specs": [
            {
              "leaf_spec": {
                "hash": "SHA256",
                "prehash_key": "NO_HASH",
                "prehash_value": "SHA256",
                "length": "VAR_PROTO",
                "prefix": "AA=="
              },
              "inner_spec": {
                "child_order": [0, 1],
                "child_size": 33,
                "min_prefix_length": 4,
                "max_prefix_length": 12,
                "empty_child": null,
                "hash": "SHA256"
              },
              "max_depth": 0,
              "min_depth": 0,
              "prehash_key_before_comparison": false
            },
            {
              "leaf_spec": {
                "hash": "SHA256",
                "prehash_key": "NO_HASH",
                "prehash_value": "SHA256",
                "length": "VAR_PROTO",
                "prefix": "AA=="
              },
              "inner_spec": {
                "child_order": [0, 1],
                "child_size": 32,
                "min_prefix_length": 1,
                "max_prefix_length": 1,
                "empty_child": null,
                "hash": "SHA256"
              },
              "max_depth": 0,
              "min_depth": 0,
              "prehash_key_before_comparison": false
            }
          ],
          "upgrade_path": ["upgrade", "upgradedIBCState"],
          "allow_update_after_expiry": true,
          "allow_update_after_misbehaviour": false
        },
        "consensus_state": {
          "timestamp": "2024-10-17T07:47:33.124389629Z",
          "root": {
            "hash": "cgIJagBEc/5lDkWS12NG5i7SSZ5hNFlDrlparFaWytc="
          },
          "next_validators_hash": "632730A03DEF630F77B61DF4092629007AE020B789713158FABCB104962FA54F"
        },
        "initial_val_set": [
          {
            "pub_key": {
              "ed25519": "RrclQz9bIhkIy/gfL485g3PYMeiIku4qeo495787X10="
            },
            "power": "500"
          }
        ]
      },
      "new_chain": true
    }
```
