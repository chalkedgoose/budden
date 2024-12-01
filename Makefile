# This Makefile is for basic tool setup in our repo.
# We use the Magefile for running our tasks.

.PHONY: install-tools

install-tools:
	@if [ "$(shell uname)" != "Darwin" ]; then \
		echo "Error: This script is intended for macOS only."; \
		exit 1; \
	fi
	@if ! command -v brew >/dev/null 2>&1; then \
		echo "Error: Homebrew is not installed. Please install Homebrew first."; \
		exit 1; \
	fi
	brew install mage
	brew install goose