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
  - dir: A string with the path in which to run the program.
    Optional, defaults to the current directory.
