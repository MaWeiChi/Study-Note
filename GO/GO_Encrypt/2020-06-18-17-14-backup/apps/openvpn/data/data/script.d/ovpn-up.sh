#!/bin/bash

data="{\"network\": {\"netmask\": \"$route_netmask_1\", \"gateway\": \"$route_gateway_1\"}, \"status\": true}"

ip link set $dev up

ovpnd api -r "http://127.0.0.1:58000/openvpn/status" -d "$data"

exit 0
