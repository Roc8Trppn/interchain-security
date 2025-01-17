#!/bin/bash

BINARY="./build/interchain-security-pd"

# Define wallet names
SENDER="public"
RECEIVER="alm"

AMOUNT="1000000000stake"
CHAIN_ID="provider-chain"
FEES="5000stake"
GAS="200000"
SENDER_ADDRESS=$($BINARY keys show "$SENDER" -a)
RECEIVER_ADDRESS=$($BINARY keys show "$RECEIVER" -a)

# Check if addresses were retrieved
if [ -z "$SENDER_ADDRESS" ] || [ -z "$RECEIVER_ADDRESS" ]; then
  echo "Error: Unable to retrieve addresses for $SENDER or $RECEIVER."
  exit 1
fi

echo "sender address $SENDER_ADDRESS" 
echo "reciever address $RECEIVER_ADDRESS" 

# Automate the transaction //
echo "Sending $AMOUNT from $SENDER to $RECEIVER..."
$BINARY tx bank send "$SENDER_ADDRESS" "$RECEIVER_ADDRESS" "$AMOUNT" --fees "$FEES" --gas "$GAS" --chain-id "$CHAIN_ID" -y

# Verify transaction submission
if [ $? -eq 0 ]; then
  echo "Transaction successfully submitted!"
else
  echo "Transaction failed!"
fi
