test:
	go list ./... | grep -v /vendor/ | go test -v

bench:
		go list ./... | grep -v /vendor/ | go test -v -bench=.
