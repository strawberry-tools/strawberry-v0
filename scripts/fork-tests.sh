#!/usr/bin/env bash

# These are a few simple tests to make sure the repository directory is as we
# expect. This is mostly used to prevent errors when merging upstream changes.
# This should be ran from the root of the repository.

failed=""

# Make sure the ./docs directory doesn't exist
if [[ -d "./docs" ]];then
	echo "Error: The directory './docs' shouldn't exist."
	failed="true"
fi

# Make sure that we pulled an actual release and not a dev version
if grep -q "DEV" ./src/common/hugo/version_current.go;then
	echo "Error: It looks like a dev version of Hugo was pulled."
	failed="true"
fi

# Make sure we're not using packages from github.com/gohugoio/hugo
if grep -q "github.com\/gohugoio\/hugo" ./src/go.mod;then
	echo "Error: It looks like a package from github.com/gohugoio/hugo is being used."
	failed="true"
fi

# Catch any files with the "extended" build tag.
# We remove it in Strawberry as we only do the "extended" release.
if grep -rq "+build extended" ./src;then
	echo "Error: the extended build tag is in use."
	failed="true"
fi

# Catch any files with the "!extended" build tag.
# We remove it in Strawberry as we only do the "extended" release.
if grep -rq "+build !extended" ./src;then
	echo "Error: the !extended build tag is in use."
	failed="true"
fi

if [[ $failed != "" ]];then
	echo "One or more tests failed."
	exit 1
else
	echo "All the \"Fork Tests\" passed."
fi
