#!/bin/bash

# generate blockchain
rivinecg generate blockchain

# keep our frontend icon packet
git checkout frontend/explorer/public/assets/icon.png frontend/explorer/public/assets/icon.ico

# update the vendor deps
dep ensure -v -update
