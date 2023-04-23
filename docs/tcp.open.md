check = "tcp.open"

Passes if a TCP connection can be opened.
Useful for checking many types of services,
such as databases and caches.
There is a 9 second timeout.

  - address: A string containing the host:port combo to connect to.
  - label: A string with a descriptive label. Optional.
