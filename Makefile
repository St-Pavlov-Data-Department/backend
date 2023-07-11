TEST := MallocNanoZone=0 go test -count=1 -race -cover -coverprofile=coverage.out

test:
	$(TEST) $$(go list ./... | grep -v vendor/ | grep -v controller)

build:
	go build -ldflags "$(GO_LDFLAGS)" -a -tags "netgo osusergo" -installsuffix netgo -o stpavlov-backend
