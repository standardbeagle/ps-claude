.PHONY: build clean install

# Load environment variables from .env file
include .env
export

build:
	GOOS=windows GOARCH=amd64 go build -o ps-claude.exe

clean:
	rm -f ps-claude.exe

install: build
	@# Convert Windows path to WSL format
	@WSL_PATH=$$(echo "$(WINDOWS_USERS_PATH)" | sed -E 's/^([A-Za-z]):/\/mnt\/\L\1/; s/\\/\//g'); \
	echo "Installing ps-claude.exe to $$WSL_PATH/$(WINDOWS_USER)/.local/bin"; \
	mkdir -p "$$WSL_PATH/$(WINDOWS_USER)/.local/bin"; \
	cp ps-claude.exe "$$WSL_PATH/$(WINDOWS_USER)/.local/bin/"; \
	echo "Installation complete!"