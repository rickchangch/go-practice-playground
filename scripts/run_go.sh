#!/bin/bash
set -e -u

SCRIPT_DIR="$( cd "$( dirname "$0" )" && pwd )"
SRC_DIR=${SCRIPT_DIR}/../src
TARGET_DIR=""

# parse parameters
while getopts t: option
do
  case $option in
    t) TARGET_DIR="$OPTARG"
    ;;
  esac
done

# run
go run ${SRC_DIR}/${TARGET_DIR}/main.go 

echo "done."