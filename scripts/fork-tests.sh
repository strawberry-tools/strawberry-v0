#!/usr/bin/env bash

# These are a few simple test to make sure the repository directory is as we expect.
# This is mostly used to prevent errors when merging upstream changes.
# This should be ran from the root of the repository.

failed=""

# Make sure the ./docs directory doesn't exist
if [[ -d "./docs" ]];then
	echo "Error: The directory './docs' shouldn't exist."
	failed="true"
fi

# Make sure that we pulled an actual release and not a dev version
if grep -q "DEV" ./common/hugo/version_current.go;then
	echo "Error: It looks like a dev version of Hugo was pulled."
	failed="true"
fi

if [[ $failed != "" ]];then
	echo "One or more tests failed."
	exit 1
else
	echo "All the \"Fork Tests\" passed."
fi
