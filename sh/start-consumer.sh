#!/bin/bash
BINARY="./build/interchain-security-pd"
CONSUMER_DATA_PATH="$HOME/.interchain-security-c/data"
CONSUMER_GENESIS_PATH="$HOME/.interchain-security-c/config/genesis.json"
CONSUMER_BINARY="./build/interchain-security-cd"

find "$CONSUMER_DATA_PATH" -type f ! -name 'priv_validator_state.json' -delete

PROVIDER_BLOCK_TIME=$($BINARY status | jq -r '.sync_info.latest_block_time')

echo "The provider block time: $PROVIDER_BLOCK_TIME"

sed -i '' "s/\"genesis_time\":.*/\"genesis_time\": \"$PROVIDER_BLOCK_TIME\",/" $CONSUMER_GENESIS_PATH

$CONSUMER_BINARY start