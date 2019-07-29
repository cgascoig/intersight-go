BASE=github.com/cgascoig/intersight-go
CLI_SRC = $(wildcard cli/*.go) $(wildcard cli/**/*.go)
SWAGGER_CLI = cli/commands/cli.go
SWAGGER_SPEC = restapi-v2.json

TEST_CMD ?= go test -v
TEST_ARGS ?= 

BUILD_CMD ?= go build -v

.PHONY: all test

# Build a binary for the current OS/arch
dist/intersight-cli: dist ${CLI_SRC} Makefile
	${BUILD_CMD} -o $@ ${BASE}/cli
	
all:  dist/intersight-cli-darwin_x64 dist/intersight-cli-linux_x64 dist/intersight-cli-win_x64.exe

dist:
	mkdir -p dist

# Build binary for all OS's
dist/intersight-cli-darwin_x64: dist ${CLI_SRC} Makefile
	GOOS=darwin GOARCH=amd64 ${BUILD_CMD} -o $@ ${BASE}/cli
dist/intersight-cli-linux_x64: dist ${CLI_SRC} Makefile
	GOOS=linux GOARCH=amd64 ${BUILD_CMD} -o $@ ${BASE}/cli
dist/intersight-cli-win_x64.exe: dist ${CLI_SRC} Makefile
	GOOS=windows GOARCH=amd64 ${BUILD_CMD} -o $@ ${BASE}/cli

$(SWAGGER_CLI): $(SWAGGER_SPEC)
	swagger generate client -f $(SWAGGER_SPEC) --skip-validation -t cli

test: # TODO