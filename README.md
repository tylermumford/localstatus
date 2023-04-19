# friendly-broccoli

Automatically monitors the important parts of your local development environment.


## Example

```
$ brock
OK  Database service
OK  Database migrations
OK  API service
OK  Env vars
!   VPN connection
OK  Redis
( exit code 1 )
```

The configuration looks like this:

```toml
# brock_checks.toml
health_checks = [
    {check = "brock.database.up", connection = "localhost:3306"},
    {check = "brock.database.migrated", tool = "flyway", directory = "~/dev/proj/db"},
    {check = "brock.http.ok", url = "http://localhost:9000/api/ping", label = "API server"},
    {check = "brock.env", variables_required = ["API_KEY", "ENVIRONMENT_MODE"]},
    {check = "brock.http.ok", url = "https://corp.private.example.com", label = "VPN connection"},
    {check = "brock.redis.up", connection = "localhost:6379"},
]
```
