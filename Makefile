test:
	go list ./... | grep -v /vendor/ | go test -v
