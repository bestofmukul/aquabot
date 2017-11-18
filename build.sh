#!/bin/bash

if [[ -z $1 ]]; then
   echo "Please provide a path to build"
   exit 1
fi

GOARM=7 GOARCH=arm GOOS=linux go build $1
