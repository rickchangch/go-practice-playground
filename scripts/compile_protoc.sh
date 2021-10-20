#!/bin/bash
set -e -u

echo "compiling proto files..."

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

# compile protoc
cd ${SRC_DIR}/${TARGET_DIR}/proto
protoc *.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative

echo "done."