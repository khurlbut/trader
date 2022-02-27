#!/bin/sh

timestamp=`date +%s000`

api_key=$1
secret_key=$2

api_url="https://api.binance.us"

signature=`echo -n "timestamp=$timestamp" | openssl dgst -sha256 -hmac $secret_key`

curl -X "GET" "$api_url/api/v3/account?timestamp=$timestamp&signature=$signature" \
     -H "X-MBX-APIKEY: $api_key"