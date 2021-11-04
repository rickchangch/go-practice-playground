#!/bin/bash
set -e -u

SCRIPT_DIR="$( cd "$( dirname "$0" )" && pwd )"
SRC_DIR=${SCRIPT_DIR}/../src
TARGET_DIR=""
VERBOSE=""

# parse parameters
while getopts "vt:" option
do
  case $option in
    v) VERBOSE="-v"
    ;;
    t) TARGET_DIR="$OPTARG"
    ;;
  esac
done

# run
go test ${VERBOSE} ${SRC_DIR}/${TARGET_DIR}/...

echo "=== script done ==="