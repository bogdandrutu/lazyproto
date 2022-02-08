export GOBIN=$(GOPATH)/bin
export PROTOBUF_ROOT=/usr/local/Cellar/protobuf/3.19.4

.PHONY: install test gen-conformance ge-testproto

install:
	go install -tags protolegacy google.golang.org/protobuf/cmd/protoc-gen-go
	go install -tags protolegacy ./cmd/protoc-gen-go_lazyproto

PROTO_FILES := $(wildcard testproto/*/*.proto)

gen-testproto: install
	for name in $(PROTO_FILES); do \
		$(PROTOBUF_ROOT)/bin/protoc \
			--go_lazyproto_out=. --plugin protoc-gen-go="${GOBIN}/protoc-gen-go_lazyproto" $${name}; \
  	done

OTEL_DOCKER_PROTOBUF ?= otel/build-protobuf:0.9.0
PROTOC := docker run --rm -u ${shell id -u} -v${PWD}:${PWD} -w${PWD} ${OTEL_DOCKER_PROTOBUF} --proto_path=${PWD}
PROTO_INCLUDES := -I/usr/include/github.com/gogo/protobuf

PROTO_GEN_GO_DIR ?= gen-gogo


# Function to execute a command. Note the empty line before endef to make sure each command
# gets executed separately instead of concatenated with previous one.
# Accepts command to execute as first parameter.
define exec-command
$(1)

endef

# Generate gRPC/Protobuf implementation for Go.
.PHONY: gen-go
gen-go:
	rm -rf ./$(PROTO_GEN_GO_DIR)
	mkdir -p ./$(PROTO_GEN_GO_DIR)
	for name in $(PROTO_FILES); do \
		$(PROTOC) $(PROTO_INCLUDES) --gogofaster_out=./$(PROTO_GEN_GO_DIR) $${name}; \
  	done
