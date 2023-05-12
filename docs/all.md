check = "command"

Experimental.
(I may decide to run this through shells in the future.)

Runs a command directly (not through a shell).
Passes only if the command's exit code is zero.
Note that the output will look like there are no quotes,
but arguments are kept intact (quoted) when passed to the program.

  - program: A string with the name of the program to run.
    If you want to run something in the current directory,
    use ./ before the file name.
    If the program is on your PATH, it will be found by its name alone.
  - args: An array of strings containing any arguments to pass.
    Optional.
  - dir: TODO.

----------

check = "const"

Simply returns the information
given when created.
Useful for debug messages
and grouping other checks.

  - okay: A bool indicating success or failure.
  - label: A string containing the message for the user.

----------

check = "env"

Passes if all the given environment variables are set.
Variables set to an empty string DO count as being set.

  - variables_required: An array of strings,
    each with an env var name.

----------

check = "file.exists"

Passes if the given file exists.
Works for both files and directories.

  - path: A string containing the path to check.

----------

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

----------

check = "http.ok"

Passes if the URL responds with 200 OK. Uses the GET method.

  - url: A string containing the URL to send an HTTP/HTTPS request to.

----------

check = "npm.install"

Passes if npm dependencies are installed.
This is implemented with the `npm ls` command,
so npm must be installed and on the PATH already.

  - package: A string with the path to a package.json file.

----------

check = "tcp.open"

Passes if a TCP connection can be opened.
Useful for checking many types of services,
such as databases and caches.
There is a 9 second timeout.

  - address: A string containing the host:port combo to connect to.
  - label: A string with a descriptive label. Optional.
