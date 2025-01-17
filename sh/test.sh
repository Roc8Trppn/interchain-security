#!/bin/bash

APP_TOML_FILE="./tomls/app.toml"
CONFIG_TOML_FILE="./tomls/config.toml"
CLIENT_TOML_FILE="./tomls/client.toml"

DEST_APP_TOML="$HOME/.interchain-security-c/config/app.toml"
DEST_CONFIG_TOML="$HOME/.interchain-security-c/config/config.toml"
DEST_CLIENT_TOML="$HOME/.interchain-security-c/config/client.toml"

# Check if the source file exists
if [[ ! -f "$APP_TOML_FILE" ]]; then
    echo "Source file does not exist: $APP_TOML_FILE"
    exit 1
fi

# Check if the destination directory exists
DEST_DIR=$(dirname "$DEST_APP_TOML")
if [[ ! -d "$DEST_DIR" ]]; then
    echo "Destination directory does not exist: $DEST_DIR"
    exit 1
fi

cp "$APP_TOML_FILE" "$DEST_APP_TOML"
echo "Replaced app.toml successfully"

# Check if the source file exists
if [[ ! -f "$CONFIG_TOML_FILE" ]]; then
    echo "Source file does not exist: $CONFIG_TOML_FILE"
    exit 1
fi

# Check if the destination directory exists
DEST_DIR=$(dirname "$DEST_CONFIG_TOML")
if [[ ! -d "$DEST_DIR" ]]; then
    echo "Destination directory does not exist: $DEST_DIR"
    exit 1
fi

cp "$CONFIG_TOML_FILE" "$DEST_CONFIG_TOML"
echo "Replaced config.toml successfully"

# Check if the source file exists
if [[ ! -f "$CLIENT_TOML_FILE" ]]; then
    echo "Source file does not exist: $CLIENT_TOML_FILE"
    exit 1
fi

# Check if the destination directory exists
DEST_DIR=$(dirname "$DEST_CLIENT_TOML")
if [[ ! -d "$DEST_DIR" ]]; then
    echo "Destination directory does not exist: $DEST_DIR"
    exit 1
fi

cp "$CLIENT_TOML_FILE" "$DEST_CLIENT_TOML"
echo "Replaced client.toml successfully"
