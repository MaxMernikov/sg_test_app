# migration
https://github.com/pressly/goose

go get -u github.com/pressly/goose/v3/cmd/goose
export GOOSE_DRIVER=clickhouse
export GOOSE_DBSTRING=tcp://localhost:9000?username=&compress=true&database=saygames_test

goose -dir ./migrations status
goose -dir ./migrations up

