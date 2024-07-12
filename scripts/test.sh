#!/bin/bash

# Generate a random URL
generate_random_url() {
  echo "http://$(tr -dc 'a-z0-9' < /dev/urandom | head -c 10).com"
}

# Create a temporary file to hold the JSON body
touch payload.json

# Generate a random URL and write it to the JSON body
generate_body() {
  local url=$(generate_random_url)
  echo "{\"url\": \"$url\"}" > payload.json
}

generate_body

# Loop to make 'n' requests
for ((i=0; i<$(shuf -i 1-1 -n 1) ; i++)); do
  generate_body
  hey -h2 -m POST -n $(shuf -i 100-100 -n 1) -c $(shuf -i 10-10 -n 1) -D payload.json -H "Content-Type: application/json" http://localhost:8000/short
done

# Clean up temporary file
rm payload.json