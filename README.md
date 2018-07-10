## Building

Go Plugins are currently only supported on Linux / MacOS

    `make`


## Testing

Create docker container

    `docker run --name testmysql -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -p 3306:3306  -d mysql`

    `./go-plugin-sql`
