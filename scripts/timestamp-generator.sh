#!/usr/bin/env bash

# Create a file whose date contents only changes every X number of days where
# X is "interval".

interval=4

day=$( date +"%d")
month=$( date +"%m")
remainder=$(( day  % interval ))
timestamp=$month/$(( day - remainder ))

echo $timestamp > ./GO_CACHE_TIMESTAMP
