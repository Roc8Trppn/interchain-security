#!/bin/bash

BINARY="./build/interchain-security-pd"
# Define wallet names
SENDER="validator"
CHAIN_ID="provider-chain"
FEES="5000stake"
GAS="200000"
SENDER_ADDRESS=$($BINARY keys show "$SENDER" -a)

JSON_FILE_PATH="./chain_proposal.json"

# Check if addresses were retrieved
# if [ -z "$SENDER_ADDRESS" ] || [ -z "$RECEIVER_ADDRESS" ]; then
#   echo "Error: Unable to retrieve addresses for $SENDER or $RECEIVER."
#   exit 1
# fi

echo "Proposal by $SENDER with  address $SENDER_ADDRESS" 

# Automate the proposal
$BINARY tx gov submit-proposal "$JSON_FILE_PATH" --from="$SENDER" --fees "$FEES" --gas "$GAS" --chain-id "$CHAIN_ID"

# Verify transaction submission
# if [ $? -eq 0 ]; then
#   echo "Proposal successfully submitted!"
# else
#   echo "Proposal submit failed!"
# fi
