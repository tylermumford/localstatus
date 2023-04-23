check = "git.branch"

Passes if the current branch
has all of the commits of the remote base branch.
(For example, if a feature branch
has all of the commits of origin/main.)

Runs `git fetch` and `git log`.
Git must already be on PATH.

  - dir: A string containing the path to the Git directory.
  - base: A string containing the name of the base branch.
    (TODO: Defaults to "origin/main".)
