.PHONY: build clean

build:
	GOOS=windows GOARCH=amd64 go build -o ps-claude.exe

clean:
	rm -f ps-claude.exe