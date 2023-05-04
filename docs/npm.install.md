check = "npm.install"

Passes if npm dependencies are installed.
This is implemented with the `npm ls` command,
so npm must be installed and on the PATH already.

  - package: A string with the path to a package.json file.
