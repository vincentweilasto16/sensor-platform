#!/bin/sh
# wait-for-it.sh
# Usage: ./wait-for-it.sh host:port -- command args

set -e

hostport="$1"
shift
cmd="$@"

host=$(echo $hostport | cut -d: -f1)
port=$(echo $hostport | cut -d: -f2)

echo "⏳ Waiting for $host:$port..."

until nc -z "$host" "$port"; do
    sleep 2
done

echo "✅ $host:$port is available. Starting command..."
exec bash -c "$cmd"
