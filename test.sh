export TESTING=true
go test $(go list ./... | grep -v /api/ | grep -v /domain | grep -v /interfaces) -coverprofile coverage.out
go tool cover -html=coverage.out