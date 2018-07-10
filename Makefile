all: go-plugin-sql plugins/sql-provider.so


go-plugin-sql: *.go
	go build -o go-plugin-sql .


plugins/sql-provider.so: plugins/sql-provider/*.go
	go build -buildmode=plugin -o plugins/sql-provider.so ./plugins/sql-provider/

