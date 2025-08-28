#!/bin/bash
set -e

# Wait for Kafka
./wait-for-it.sh kafka:9092 "/bin/true"
echo "✅ kafka:9092 is available"

# Start service
exec ./service-a