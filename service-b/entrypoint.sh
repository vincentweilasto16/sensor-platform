#!/bin/bash
set -e

# Wait for MySQL
./wait-for-it.sh mysql:3306 "/bin/true"
echo "✅ mysql:3306 is available"

# Wait for Kafka
./wait-for-it.sh kafka:9092 "/bin/true"
echo "✅ kafka:9092 is available"

# Start the service
exec ./service-b
