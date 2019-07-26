BASE=github.com/cgascoig/intersight-go
CLI_SRC = $(wildcard cli/*.go) $(wildcard cli/**/*.go)
SWAGGER_CLI = cli/commands/cli.go
SWAGGER_SPEC = restapi-v2.json

TEST_CMD ?= go test -v
TEST_ARGS ?= 

BUILD_CMD ?= go build -v

.PHONY: test all
	
all:  dist/intersight-cli

dist:
	mkdir -p dist

dist/intersight-cli: dist ${CLI_SRC} Makefile
	${BUILD_CMD} -o $@ ${BASE}/cli
	GOOS=linux GOARCH=amd64 ${BUILD_CMD} -o $@-linux_x64 ${BASE}/cli

$(SWAGGER_CLI): $(SWAGGER_SPEC)
	swagger generate client -f $(SWAGGER_SPEC) --skip-validation -t cli

test: # TODO