#!/bin/bash

# Configuration
GENESIS_PATH="$HOME/.interchain-security-p/config/genesis.json"
BINARY="./build/interchain-security-pd"
ACCOUNTS=(
  "validator"
  "alm"
  "liquidity"
  "public"
  "team"
  "advisor"
  "pre-seed"
  "seed"
  "private"
  "minting"
  "community-and-dev"
  "foundation"
)   

echo "Starting account deletions..."

# Add validator account
for account in "${ACCOUNTS[@]}"; do
  IFS=":" read -r wallet_name <<< "$account"
  echo "y" | $BINARY keys delete "$wallet_name" --keyring-backend os
done

# Backup the original genesis.json
cp "$GENESIS_PATH" "${GENESIS_PATH}.bak"

jq '.app_state.auth.accounts = [] | .app_state.bank.balances = []' \
  "$GENESIS_PATH" > tmp_genesis.json && mv tmp_genesis.json "$GENESIS_PATH"

