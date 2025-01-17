#!/bin/bash

# Prerequisits
# - already working provider chain
# - ready consumer binary

# Variables
CONSUMER_BINARY="./build/interchain-security-cd"
CONSUMER_CONFIG="$HOME/.interchain-security-c"

PROVIDER_CHAIN_ID="provider-chain"
CONSUMER_CHAIN_ID="zone_chain"
MONIKER="zone-node"

PROVIDER_BINARY="./build/interchain-security-pd"
CONSUMER_JSON="./proposals/create_consumer.json"

SIGNER_NAME="validator"

# Remove existing consumer config
set -e
rm -rf "$CONSUMER_CONFIG"
echo "Removed Consumer Config"

# Reinitialize the Consumer Chain
$CONSUMER_BINARY init "$MONIKER" --chain-id "$CONSUMER_CHAIN_ID"
echo "Consumer chain initialized successfully"

# Generate Genesis Hash and Binary Hash of the genesis.json
GENESIS_HASH=$(shasum -a 256 $CONSUMER_CONFIG/config/genesis.json | awk '{print $1}')
echo "Genesis Hash: $GENESIS_HASH"

BINARY_HASH=$(shasum -a 256 $CONSUMER_BINARY | awk '{print $1}')
echo "Binary Hash: $BINARY_HASH"

# Update Consumer Chain json  with Chain-ID
jq --arg chain_id "$CONSUMER_CHAIN_ID" \
   '.chain_id = $chain_id' \
   "$CONSUMER_JSON" > temp.json && mv temp.json "$CONSUMER_JSON"
echo "Updated $CONSUMER_JSON with Chain-ID $CONSUMER_CHAIN_ID"

SPAWN_TIME=$(date -u -v+1M +"%Y-%m-%dT%H:%M:%SZ")
echo "Spawn Time is: $SPAWN_TIME"

# Update Consumer Chain json with Hashes and Spawn Time
jq --arg genesis_hash "$GENESIS_HASH" \
    --arg binary_hash "$BINARY_HASH" \
    --arg spawn_time "$SPAWN_TIME" \
   '.initialization_parameters.genesis_hash = $genesis_hash | .initialization_parameters.binary_hash = $binary_hash | .initialization_parameters.spawn_time = $spawn_time' \
   "$CONSUMER_JSON" > temp.json && mv temp.json "$CONSUMER_JSON"
echo "Updated $CONSUMER_JSON with the new hashes."

# Exchange ports in TOML Files
APP_TOML_FILE="./tomls/app.toml"
CONFIG_TOML_FILE="./tomls/config.toml"
CLIENT_TOML_FILE="./tomls/client.toml"
DEST_APP_TOML="$HOME/.interchain-security-c/config/app.toml"
DEST_CONFIG_TOML="$HOME/.interchain-security-c/config/config.toml"
DEST_CLIENT_TOML="$HOME/.interchain-security-c/config/client.toml"
if [[ ! -f "$APP_TOML_FILE" ]]; then
    echo "Source file does not exist: $APP_TOML_FILE"
    exit 1
fi
DEST_DIR=$(dirname "$DEST_APP_TOML")
if [[ ! -d "$DEST_DIR" ]]; then
    echo "Destination directory does not exist: $DEST_DIR"
    exit 1
fi
cp "$APP_TOML_FILE" "$DEST_APP_TOML"
echo "Replaced app.toml successfully"
if [[ ! -f "$CONFIG_TOML_FILE" ]]; then
    echo "Source file does not exist: $CONFIG_TOML_FILE"
    exit 1
fi
DEST_DIR=$(dirname "$DEST_CONFIG_TOML")
if [[ ! -d "$DEST_DIR" ]]; then
    echo "Destination directory does not exist: $DEST_DIR"
    exit 1
fi
cp "$CONFIG_TOML_FILE" "$DEST_CONFIG_TOML"
echo "Replaced config.toml successfully"
if [[ ! -f "$CLIENT_TOML_FILE" ]]; then
    echo "Source file does not exist: $CLIENT_TOML_FILE"
    exit 1
fi
DEST_DIR=$(dirname "$DEST_CLIENT_TOML")
if [[ ! -d "$DEST_DIR" ]]; then
    echo "Destination directory does not exist: $DEST_DIR"
    exit 1
fi
cp "$CLIENT_TOML_FILE" "$DEST_CLIENT_TOML"
echo "Replaced client.toml successfully"

# Run CreateConsumerMessage
$PROVIDER_BINARY tx provider create-consumer $CONSUMER_JSON --from=$SIGNER_NAME --gas auto --fees 5000stake --chain-id $PROVIDER_CHAIN_ID -y
echo "Consumer creation succeeded!"
sleep 3

# Grab Consumer Id
#TODO: get consumer chain either after the result from above or with this call:
#$PROVIDER_BINARY query provider consumer-chain 
CONSUMER_ID="0"
echo "Created Consumer Chain with ID $CONSUMER_ID"

# Opt-In with Signer (Validator)
$PROVIDER_BINARY tx provider opt-in "$CONSUMER_ID" --from "$SIGNER_NAME" --gas auto --fees 5000stake -y
echo "Successfully opted-in with $SIGNER_NAME"

# Wait until Spawn Time reached
echo "Awaiting Spawn Time..."
sleep 90
echo "Constinuing script..."

# Grab Consumer Chain ccvconsumer data
CONSUMER_FILE_NAME="ccvconsumer_gen.json"
$PROVIDER_BINARY query provider consumer-genesis "$CONSUMER_ID" -o json > "$CONSUMER_FILE_NAME"
# echo "Exported ccvconsumer data."

# Update the Consumer Genesis File
jq '.app_state.ccvconsumer = input' $CONSUMER_CONFIG/config/genesis.json $CONSUMER_FILE_NAME > temp.json && mv temp.json $CONSUMER_CONFIG/config/genesis.json
echo "Replaced ccvconsumer data in consumer's gensis file."

# Validate Genesis File
# $CONSUMER_BINARY genesis validate
# echo "Genesis file is valid"

# Validate Genesis File
$CONSUMER_BINARY start

# TODO: if necessary, update genesis_time and app_hash -> debug error

# TODO: Add persistent_peer
# For now  add manually into persistens_peers: bb31ae15e86a7a1d4f70f1269a7fcfa5f3000b4d@127.0.0.1:26656

# Configure hermes?
# hermes create connection --a-chain <consumer chain ID> --a-client 07-tendermint-0 --b-client <client assigned by provider chain> 
# hermes create channel --a-chain <consumer chain ID> --a-port consumer --b-port provider --order ordered --a-connection connection-0 --channel-version 1
# hermes start