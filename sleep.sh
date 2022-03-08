#!/bin/bash

flag=true

count=0

while [ $flag = true ]; do
  sleep 2
  echo "time is $(date)"
  count=$((count + 1))
  echo "$count"
  if [ $count -gt 5 ]; then
     break
  fi
done
