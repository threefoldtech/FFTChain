#!/bin/bash

# generate blockchain
rivinecg generate blockchain --explorer=plainjs

# update the vendor deps
dep ensure -v -update
