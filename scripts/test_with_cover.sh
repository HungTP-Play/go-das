rm -rf coverage
mkdir coverage
go test ./... -coverprofile=coverage/cover.out
go tool cover -func=coverage/cover.out
go tool cover -html=coverage/cover.out