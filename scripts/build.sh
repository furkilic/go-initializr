#!/usr/bin/env bash

DIR=`dirname ${0}`
. $DIR/common.sh


$BASE_DIR/gow build -o $BASE_DIR/bin/go-initializr$extension $BASE_DIR/cmd/go-initializr/main.go