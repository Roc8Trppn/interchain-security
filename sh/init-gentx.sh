#!/bin/bash

# Prerequisits:
# -> must create build folder before (see README.md)

#TODO:
# - Replace all investor, advisor and team wallets with already created addresses
# - Exact days of lockup, vesting

# Configuration
GENESIS_PATH="$HOME/.interchain-security-p/config/genesis.json"
BINARY="./build/interchain-security-pd"
CHAIN_ID="provider-chain"
TOKEN_DENOM="stake"
TOTAL_SUPPLY="1000000000000000" 

VALIDATOR_ACCOUNT=(
  "validator:50000000000000:5000000000000"
)             

# Check if gensis file exist. Create if it doesn't exist.
if [ ! -f "$GENESIS_PATH" ]; then
  echo "Initializing node as genesis.json is missing..."
  $BINARY init "hub-node" --chain-id "$CHAIN_ID"
fi

echo "Gensis.json exists or was successfully created..."

# Function to get wallet address
get_wallet_address() {
  local wallet_name=$1
  $BINARY keys show "$wallet_name" -a
}

# Add validator account
echo "Adding validator account..."
for account in "${VALIDATOR_ACCOUNT[@]}"; do
  IFS=":" read -r wallet_name amount self_stake <<< "$account"
  
  # Create the wallet (if it doesn't already exist)
  if ! $BINARY keys show "$wallet_name" &> /dev/null; then
    $BINARY keys add "$wallet_name" --output json > /dev/null
    echo "Created wallet: $wallet_name"
  fi
  
  # Get the wallet address
  wallet_address=$(get_wallet_address "$wallet_name")
  
  # Add to genesis
  $BINARY genesis add-genesis-account "$wallet_address" "${amount}${TOKEN_DENOM}"
  echo "Added $amount $TOKEN_DENOM to $wallet_name ($wallet_address)"

  # Add gentx entry to create Validator with self-delegation
  $BINARY genesis gentx "$wallet_name" "${self_stake}${TOKEN_DENOM}" --chain-id "$CHAIN_ID"
  echo "Created gentx-file successully."

  # Collect gentx information
  $BINARY genesis collect-gentxs
  echo "Collected gentx-file successully."

done
