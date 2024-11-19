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
TOTAL_SUPPLY="1000000000000000" # Total supply in smallest units (e.g., 1 billion with 6 decimals)         

# Unvested Accounts (name:amount)
GENESIS_ACCOUNTS=(
  "alm:100000000000000"
  "liquidity:50000000000000"
  "public:50000000000000"
)

# Define vesting accounts (name:amount:lockup in days : vesting in days )
VESTING_ACCOUNTS=(
  "team:130000000000000:730:1460"
  "advisor:50000000000000:730:1460"
  "pre-seed:10000000000000:365:730"
  "seed:40000000000000:365:730"
  "private:50000000000000:548:730"
  "community-and-dev:70000000000000:0:2190"
  "foundation:100000000000000:365:2190"
)

# Function to get wallet address
get_wallet_address() {
  local wallet_name=$1
  $BINARY keys show "$wallet_name" -a
}

# Validate Genesis and Exit in case of invalid genesis.json
echo "Validating genesis file..."
$BINARY genesis validate
if [ $? -eq 0 ]; then
  echo "Genesis file is valid!"
else
  echo "Genesis file validation failed!"
  exit 1
fi

# Add genesis accounts
echo "Adding genesis accounts..."
for account in "${GENESIS_ACCOUNTS[@]}"; do
  IFS=":" read -r wallet_name amount <<< "$account"
  
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
done

# Current Unix time
CURRENT_TIME=$(date +%s)

# Add vesting accounts
echo "Adding vesting accounts..."
for vesting_account in "${VESTING_ACCOUNTS[@]}"; do
  IFS=":" read -r wallet_name amount lockup vesting <<< "$vesting_account"

  # Create the wallet (if it doesn't already exist)
  if ! $BINARY keys show "$wallet_name" &> /dev/null; then
    $BINARY keys add "$wallet_name" --output json > /dev/null
    echo "Created wallet: $wallet_name"
  fi

  # Get the wallet address
  wallet_address=$(get_wallet_address "$wallet_name")

  # Get start time and end time of vesting
  LOCKUP_DURATION=$((lockup * 24 * 60 * 60)) # Convert lockup days to seconds
  VESTING_DURATION=$((vesting * 24 * 60 * 60)) # 1 year in seconds

  # Calculate dynamic start and end times
  start_time=$((CURRENT_TIME + LOCKUP_DURATION))
  end_time=$((CURRENT_TIME + LOCKUP_DURATION + VESTING_DURATION))

  # Add to balances
  $BINARY genesis add-genesis-account "$wallet_address" \
  "${amount}${TOKEN_DENOM}" \
  --vesting-amount "${amount}${TOKEN_DENOM}" \
  --vesting-start-time "$start_time" \
  --vesting-end-time "$end_time"
  echo "Added vesting account: $wallet_name ($wallet_address) with $amount$TOKEN_DENOM vesting from $start_time to $end_time"
done

# Validate genesis file
echo "Validating genesis file..."
$BINARY genesis validate-genesis
if [ $? -eq 0 ]; then
  echo "Genesis file is valid!"
else
  echo "Genesis file validation failed!"
  exit 1
fi

echo "Genesis preparation complete."
