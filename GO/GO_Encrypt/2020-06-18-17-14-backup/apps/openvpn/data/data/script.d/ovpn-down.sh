#!/bin/bash

data="{\"network\": {\"netmask\": \"\", \"gateway\": \"\", \"remoteIp\": \"\", \"localIp\": \"\"}, \"status\": false}"

ovpnd api -r "http://127.0.0.1:58000/openvpn/status" -d "$data"

# If $dev not exist then exit.
cat /proc/net/dev | grep $dev
if [ $? -ne 0 ]; then
    exit
fi

ip link set $dev down

exit 0
