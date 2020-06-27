# Maintaining Gotham

This is documentation for maintainers of Gotham.
This doc covers best practices and procedures on what to do for certain maintaining tasks.


## Making Releases

### Major & Minor Releases

Checklist for major/minor releases:

1. Make sure that the milestone is complete.
1. Create a new branch called `release-vx.y.0`.
1. Edit `common/hugo/version-gotham.go`. Make sure that the version number is correct and "dev" is removed.
1. Create a full commit message: `git commit` (no `-m`)
1. The commit title should have the format: `Release: Gotham vx.y.0`
1. The first paragraph should be a friendly message summarizing the release.
1. If appropriate, a list of highlights from the release should be done next.
1. The final paragraph/line should be the output of that version's `gotham version`
1. Push branch up and open PR.
1. Add the PR to the current milestone.
1. After PR is merged, tag the commit `vx.y.0`.
1. Push tag up, which will kick off the actual release.
1. Close the current milestone.
1. Create a new branch called `dev-vx.y.0` which should target what is likely going to be the next release version.
1. Edit `common/hugo/version-gotham.go`. Make sure that the version number is correct and "dev" is added back.
1. Commit this with the message: `Open dev for next release.`.
1. Push up and open PR. Once this PR is merged, development is ready for the next version. In the meantime, the current release is likely close to being done.

### Patch Releases

Checklist for patch releases:

1. If it doesn't exists already, create a branch called `release-vX.Y.x-patches`. 'X' and 'Y' should be replaced with the appropriate major and minor version numbers. The 'x' is a literal 'x'.
1. Edit `common/hugo/version-gotham.go`. Make sure that the version number is correct and "dev" is removed.
1. Cherry pick the relevant commits from the `master` branch.
1. Create a full commit message: `git commit` (no `-m`)
1. The commit title should have the format: `Release: Gotham vx.y.z`
1. The first paragraph should explain why the patch release is necessary.
1. If there's more than one change, a list of bug fixes and security updates should be done next.
1. The final paragraph/line should be the output of that version's `gotham version`
1. Push branch up and confirm that CircleCI passes.
1. After tests pass, tag the commit `vx.y.z`.
1. Push tag up, which will kick off the actual release.
1. keep this branch around for future patch releases on this minor version.
