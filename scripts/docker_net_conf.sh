#!/bin/bash

GMWEATHER_IP_SERACH=207.8.4.1/24
GMWEATHER_IP=207.8.4.0/24

OUTADDRSHOW=$(ip addr show | grep $GMWEATHER_IP_SERACH)

if [ ! -z "$OUTADDRSHOW" -a "$OUTADDRSHOW" != " " ]; then
    echo "network already created, initializing..."
else
    docker network create --subnet $GMWEATHER_IP gmweather_net_dev
fi
