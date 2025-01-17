#!/bin/bash
CONSUMER_DATA_PATH="$HOME/.interchain-security-c/data"
CONSUMER_BINARY="./build/interchain-security-cd"

find "$CONSUMER_DATA_PATH" -type f ! -name 'priv_validator_state.json' -delete

$CONSUMER_BINARY start