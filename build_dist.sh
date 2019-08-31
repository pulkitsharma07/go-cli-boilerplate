# This script is supposed to run inside a docker container (/ Alpine linux)
#! /bin/sh

# Get the current tag (version number)
VERSION=`git describe --abbrev=0`
# Get the latest commit-id
COMMIT=`git rev-parse --short HEAD`

# Modify these accordingly
CLI_BINARY_BASE_NAME="cli"
PACKAGE="github.com/pulkitsharma07/go-cli-boilerplate"
VERSION_BUILD_VARIABLE="$PACKAGE/cmd.Version"
COMMIT_BUILD_VARIABLE="$PACKAGE/cmd.Commit"

BUILD_DIRECTORY="dist"

# Helper function to build the binary for a particular OS
function build() {
  os=$1

  echo "Compiling for $os"
  BINARY_FILE=$CLI_BINARY_BASE_NAME

  if [ $os = "darwin" ]; then
    BINARY_FILE="$BINARY_FILE-darwin"
  fi

  if [ $os = "windows" ]; then
    BINARY_FILE="$BINARY_FILE.exe"
  fi

  GOOS=$os go build \
          --ldflags "-s -w \
          -X ${COMMIT_BUILD_VARIABLE}=${COMMIT} \
          -X ${VERSION_BUILD_VARIABLE}=${VERSION}"\
          -o $BUILD_DIRECTORY/$BINARY_FILE

  if [ $? -ne 0 ]; then
    echo "Build Failed !"
    exit $?
  fi

  md5sum $BUILD_DIRECTORY/$BINARY_FILE | cut -d' ' -f 1 > $BUILD_DIRECTORY/$BINARY_FILE.md5

  echo "Done"
}

# Build Steps start from here..
echo "Initiating build for Version: $VERSION and Commit: $COMMIT"
make test
mkdir $BUILD_DIRECTORY
build "linux"
build "darwin"
build "windows"
