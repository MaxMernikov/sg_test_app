# DB Migration
https://github.com/pressly/goose

```
go get -u github.com/pressly/goose/v3/cmd/goose
export GOOSE_DRIVER=clickhouse
export GOOSE_DBSTRING=tcp://localhost:9000?username=&compress=true&database=saygames_db

goose -dir ./migrations status
goose -dir ./migrations up
```

# Run
```
go run .

or

GOOS=linux GOARCH=amd64 go build
./saygames 2>> logfile.log
```

# Test
```
go test

or

curl -X POST -d '{"client_time":"2021-09-16 02:40:16","device_id":"dbe4750c-4fce-42f5-a974-ae2c550a9e2f","device_os":"iOS 13.5.1","session":"HlLYViMxPkdqAGEH","sequence":1,"event":"app_stop","param_int":-64,"param_str":"Chillwave health."}' http://localhost:8090

or

ali --rate=1000 --duration=10s --method=POST --body-file=./test_request.txt http://localhost:8090/
```

# Debug
```
go tool pprof http://127.0.0.1:8090/debug/pprof/profile
```

# Docker
```
docker run -it --rm --link some-clickhouse-server yandex/clickhouse-client --host some-clickhouse-server
```
