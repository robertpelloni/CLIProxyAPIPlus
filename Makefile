.PHONY: build test generate-ui ui-all

# Main build target
build: generate-ui
	go build -o cli-proxy-api ./cmd/server

# Run tests
test:
	go test ./...

# Build the UI
generate-ui:
	cd ui && npm install && npm run build
	mkdir -p internal/managementasset/ui_build
	cp ui/dist/index.html internal/managementasset/ui_build/management.html

# Shortcut
ui-all: generate-ui build
