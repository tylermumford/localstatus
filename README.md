# LocalStatus

Automatically monitors the important parts of your local development environment. Not in a fancy way.
Just in a fifteen-year-old-Honda-Civic kind of way.

## Example

```
$ localstatus
    6 checks to run...
OK  Database service
OK  Database migrations
OK  API service
OK  Env vars
!   VPN connection
OK  Redis
( exit code 1 )

$ localstatus --watch
    6 checks to run, every 3 minutes...
    (Enter to re-run, Ctrl-C to stop)
```

The configuration looks like this:

```toml
# ~/.config/localstatus.toml
checks = [
    {check = "tcp.open", address = "localhost:3306", label = "MySQL"},
    {check = "tcp.open", address = "localhost:6379", label = "Redis"},
    {check = "database.migrated", tool = "flyway", directory = "~/dev/proj/db"},
    {check = "http.ok", url = "http://localhost:9000/api/ping", label = "API server"},
    {check = "env", variables_required = ["API_KEY", "ENVIRONMENT_MODE"]},
    {check = "http.ok", url = "https://corp.private.example.com", label = "VPN connection"},
]
```

## Trivia

- This project was almost named Brock, because GitHub generated a repo name of "friendly-broccoli."
- This project was almost named Checkmate, generated by ChatGPT.
- This project was almost named ldca, for Local Development is Cool Again.
- This project was almost named Local Host Checker, for the cool acronym LHC.
