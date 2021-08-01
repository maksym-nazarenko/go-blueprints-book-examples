PROJECT_ROOT = $(shell git rev-parse --show-toplevel)

build-chat:
	@go build -o $(PROJECT_ROOT)/cmd/chat/chat $(PROJECT_ROOT)/cmd/chat

run-chat: build-chat
	TEMPLATES_DIR=$(PROJECT_ROOT)/cmd/chat/templates \
		$(PROJECT_ROOT)/cmd/chat/chat

test:
	@go test ./...
