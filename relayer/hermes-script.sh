#!/bin/bash

# CONSUMER_CLIENT_ID is created on CONSUMER upon genesis
CONSUMER_CLIENT_ID="0"
CONSUMER_CHAIN_ID="zone_chain"

# PROVIDER_CLIENT_ID is created on PROVIDER upon CONSUMER spawn time: gaiad q provider list-consumer-chains
PROVIDER_CLIENT_ID=""
PROVIDER_CHAIN_ID="provider-chain"

CONFIG=$1
if [ -z "$CONFIG" ]; then 
    CONFIG=$HOME/.hermes/config.toml
fi
if [ ! -f "$CONFIG" ]; then
    echo "no config file found at $CONFIG"
    exit 1
fi

output=$(hermes --json --config $CONFIG create connection --a-chain $CONSUMER_CHAIN_ID --a-client $CONSUMER_CLIENT_ID --b-client $PROVIDER_CLIENT_ID | tee /dev/tty)
json_output=$(echo "$output" | grep 'result')
a_side_connection_id=$(echo "$json_output" | jq -r '.result.a_side.connection_id')
output=$(hermes --json --config $CONFIG create channel --a-chain $CONSUMER_CHAIN_ID --a-port consumer --b-port provider --order ordered --a-connection $a_side_connection_id --channel-version 1 | tee /dev/tty)
json_output=$(echo "$output" | grep 'result')
echo "---- DONE ----"
echo "$json_output" | jq