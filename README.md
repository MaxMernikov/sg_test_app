# migration
https://github.com/pressly/goose

go get -u github.com/pressly/goose/v3/cmd/goose
export GOOSE_DRIVER=clickhouse
export GOOSE_DBSTRING=tcp://localhost:9000?username=&compress=true&database=saygames_test

goose -dir ./migrations status
goose -dir ./migrations up


ali --rate=1 --duration=1s --method=POST --body-file=./test_request.txt http://localhost:8090/

go run .
go tool pprof http://127.0.0.1:8090/debug/pprof/profile

go test

# DOCKER CLI
docker run -it --rm --link some-clickhouse-server yandex/clickhouse-client --host some-clickhouse-server