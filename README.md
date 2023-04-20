# friendly-broccoli

Automatically monitors the important parts of your local development environment.


## Example

```
$ brock
    6 checks to run...
OK  Database service
OK  Database migrations
OK  API service
OK  Env vars
!   VPN connection
OK  Redis
( exit code 1 )

$ brock --watch
    6 checks to run, every 3 minutes...
    (Enter to re-run, Ctrl-C to stop)
```

The configuration looks like this:

```toml
# ~/brock_checks.toml
checks = [
    {check = "brock.tcp.open", address = "localhost:3306", label = "MySQL"},
    {check = "brock.tcp.open", address = "localhost:6379", label = "Redis"},
    {check = "brock.database.migrated", tool = "flyway", directory = "~/dev/proj/db"},
    {check = "brock.http.ok", url = "http://localhost:9000/api/ping", label = "API server"},
    {check = "brock.env", variables_required = ["API_KEY", "ENVIRONMENT_MODE"]},
    {check = "brock.http.ok", url = "https://corp.private.example.com", label = "VPN connection"},
]
```
