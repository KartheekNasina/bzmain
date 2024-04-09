#!/bin/sh

# Timeout & interval.
TIMEOUT=30
INTERVAL=5

echo "Waiting for Redis to be available..."

# Loop until the Redis port is accessible or we timeout.
elapsed=0
while true; do
    if [ $elapsed -ge $TIMEOUT ]; then
        echo "Timed out waiting for Redis to be available."
        exit 1
    fi
    
    # Try to connect to Redis using netcat.
    if nc -z -w 1 redis-master 6379; then
        echo "Redis is available!"
        break
    fi
    
    # Wait for the interval period.
    sleep $INTERVAL
    elapsed=$(($elapsed+$INTERVAL))
done
