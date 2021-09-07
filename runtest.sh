rm -rf utils/*.log

go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

rm -rf utils/*.log