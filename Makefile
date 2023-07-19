TEST := MallocNanoZone=0 go test -count=1 -race -cover -coverprofile=coverage.out

# -------- Build info
GitCommitHash := $(shell git rev-parse HEAD)
GoVersion := $(shell go version)

GO_LDFLAGS ?= -s \
	-X $(REPO_PATH)/buildinfo.GitCommitHash=$(GitCommitHash) \
	-X '$(REPO_PATH)/buildinfo.GoVersion=$(GoVersion)' \

GO_GCFLAGS ?= -l=4

test:
	$(TEST) $$(go list ./... | grep -v vendor/ | grep -v controller)

build:
	CGO_ENABLED=0 \
		go build \
		-gcflags "$(GO_GCFLAGS)" \
		-ldflags "$(GO_LDFLAGS)" \
		-a -tags "netgo osusergo" -installsuffix netgo \
		-o stpavlov-backend
