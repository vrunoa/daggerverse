#!/bin/sh
coverfile=$1
threshold=$2
if [[ -z $coverfile ]]; then
  echo "Please provide a coverfile"
  exit 1
fi
if [[ -z $threshold ]]; then
  echo "Please provide a threshold"
  exit 1
fi

go tool cover -func=$coverfile
out=$(go tool cover -func $coverfile | grep "total:" | awk '{print $3}' | grep -Eo '^\d*')
if [[ $out -lt $threshold ]]; then
    echo "coverage below threshold $threshold%"
    exit 1
fi

