# Strawberry Development Information

*This is a temporary file. Once the Development section of Strawberry Docs is up, this information should be moved there.*

## Cloning

Cloning is pretty much straight forward however if you want to be able to pull into Hugo releases or other Hugo tasks, there's a few extra steps.

1. First clone as you would normally: `git clone https://github.com/strawberryssg/strawberry-v0.git`
1. `cd strawberry-v0`
1. Then add the Hugo repository as a remote. We prefer the remote name "upstream" but you can use what you like: `git remote add upstream https://github.com/gohugoio/hugo.git`
1. We want to tell git fetch to not pull tags from this upstream. In the git config for this repo, set "remote.<name>.tagopt" where <name> is the Hugo upstream namae you used: `remote.upsteam.tagopt=--no-tags`
1. Then we can run git fetch: `git fetch upstream`


## Dependencies

- g++ is needed. On Ubuntu, this can be installed via the "build-essential" package.


## Notes

- **RAM usage** - Running the full test suite (included RACE tests) can consume more than 4GB of RAM.
