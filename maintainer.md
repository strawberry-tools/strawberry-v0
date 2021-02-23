# Maintaining Strawberry

This is documentation for maintainers of Strawberry.
This doc covers best practices and procedures on what to do for certain maintaining tasks.


## Making Releases

### Major & Minor Releases

Checklist for major/minor releases:

1. Make sure that the milestone is complete.
1. Create a new branch called `release-vx.y.0`.
1. Edit `src/common/hugo/version_strawberry.go`. Make sure that the version number is correct and "dev" is removed.
1. Create a full commit message: `git commit` (no `-m`)
1. The commit title should have the format: `Release: Strawberry vx.y.0`
1. The first paragraph should be a friendly message summarizing the release.
1. If appropriate, a list of highlights from the release should be done next.
1. The final paragraph/line should be the output of that version's `strawberry version`
1. Push branch up and open PR.
1. Add the PR to the current milestone.
1. After PR is merged, tag the commit `vx.y.0`.
1. Push tag up, which will kick off the actual release.
1. Close the current milestone.
1. Create a new branch called `dev-vx.y.0` which should target what is likely going to be the next release version.
1. Edit `src/common/hugo/version_strawberry.go`. Make sure that the version number is correct and "dev" is added back.
1. Commit this with the message: `Open dev for next release.`.
1. Push up and open PR. Once this PR is merged, development is ready for the next version. In the meantime, the current release is likely close to being done.

### Patch Releases

Checklist for patch releases:

1. If it doesn't exists already, create a branch called `release-vX.Y.x-patches`. 'X' and 'Y' should be replaced with the appropriate major and minor version numbers. The 'x' is a literal 'x'.
1. Edit `src/common/hugo/version_strawberry.go`. Make sure that the version number is correct and "dev" is removed.
1. Cherry pick the relevant commits from the `master` branch.
1. Create a full commit message: `git commit` (no `-m`)
1. The commit title should have the format: `Release: Strawberry vx.y.z`
1. The first paragraph should explain why the patch release is necessary.
1. If there's more than one change, a list of bug fixes and security updates should be done next.
1. The final paragraph/line should be the output of that version's `strawberry version`
1. Push branch up and confirm that CircleCI passes.
1. After tests pass, tag the commit `vx.y.z`.
1. Push tag up, which will kick off the actual release.
1. keep this branch around for future patch releases on this minor version.


## Pulling Upstream

Pulling changes from upstream is a simple but sometimes tedious process.
By default, the Strawberry Team only pulls in upstream changes when they do a release.
The SemVer bump Strawberry needs to have should be a combination of the changes within Strawberry, if any, and the type of release by Hugo.
Whichever has the highest priority between patch, minor, and major, that's what you go with.

1. Make sure `trunk` is up-to-date` and create a new branch with the Hugo version:
    ```bash
	git checkout trunk && git pull && git checkout -b hugo-vx.y.z
	```
1. Fetch the upstream remote so that git is aware of the upstream changes:
    ```bash
	git fetch upstream
	```
1. Merge in the release. To do this, find the commit has for the release tag and use it here. There is a 98% chance there will be merge conflicts. That's no problem:
    ```bash
	git merge <commit-hash>
	```
1. If you have merge conflicts, fix them, one by one. Some tips:
    - `git add` each file **as you fix it**. This ensures that whenever you run `git status`, it's an accurate reflection of the remaining work.
    - if a file is added to or modified in `./docs`, delete it with `git rm -r ./docs/`. Unlike upstream, we don't keep user docs in this repo.
    - Often you'll see conflicts in the Go imports. When this happens, an import was added/modified/removed. If one was added or modified, you also need to make sure you correct the important path if it's a module within Hugo's codebase.
1. Make sure that `mage -v test` and `mage -v check` are passing.
1. Confirm that `src/common/hugo/version_current.go` contains the Hugo version you expect.
1. Make sure to run `go mod tidy` from within the `src` directory.
1. With all conflicts resolved, run `git commit`. In the text editor, use the commit message "Upstream: Pull in changes from Hugo vx.y.z" with the appropriate version number in place.
1. Push it up and open a PR with the title being the same of the commit message. If this PR is likely to be included in the next scheduled Strawberry minor release, include it in the milestone. If a patch release is expected, there's no need.
1. When the PR is ready to be merged, temporarily turn on "merge commits" for the repo, and use a merge commit to merge the PR. Then turn off the setting again.
