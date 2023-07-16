TEST := MallocNanoZone=0 go test -count=1 -race -cover -coverprofile=coverage.out

# -------- Build info
GitCommitHash := $(shell git rev-parse HEAD)
GoVersion := $(shell go version)

GO_LDFLAGS ?= -s \
	-X $(REPO_PATH)/buildinfo.GitCommitHash=$(GitCommitHash) \
	-X '$(REPO_PATH)/buildinfo.GoVersion=$(GoVersion)' \

test:
	$(TEST) $$(go list ./... | grep -v vendor/ | grep -v controller)

build:
	go build -ldflags "$(GO_LDFLAGS)" -a -tags "netgo osusergo" -installsuffix netgo -o stpavlov-backend
