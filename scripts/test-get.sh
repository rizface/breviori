#!/bin/bash

# Generate a random URL
generate_random_url() {
  echo "http://$(tr -dc 'a-z0-9' < /dev/urandom | head -c 10).com"
}

# Loop to make 'n' requests
for ((i=0; i<$(shuf -i 1000-5000 -n 1) ; i++)); do
  generate_body
  hey -h2 -m GET -n 1 -c 1 -H "Content-Type: application/json" http://localhost:8000/redirect/$(generate_random_url)
done