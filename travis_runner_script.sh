# This script is invoked from .travis.yml
#! /bin/sh

# Need to build binaries if build is tagged (i.e. there is tag associated with the current commit)
if [[ $TRAVIS_TAG != "" ]]; then
  make build
else # If build is not triggered from a tag commit, we do not need to build the binaries.
  make test
fi
