PROJECT_ROOT = $(shell git rev-parse --show-toplevel)
build-chat:
	@go build -o ./cmd/chat/chat ./cmd/chat
