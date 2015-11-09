#!/bin/bash

echo "[" > json.txt
for num in $(seq 1 403); do curl https://xkcd.com/$num/info.0.json >> json.txt && echo "," >> json.txt; done
for num in $(seq 405 1599); do curl https://xkcd.com/$num/info.0.json >> json.txt && echo "," >> json.txt; done
curl https://xkcd.com/1600/info.0.json >> json.txt
echo "]" >> json.txt
