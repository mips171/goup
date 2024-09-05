# Makefile for goup

BINARY = goup
INSTALL_DIR = /usr/local/bin

build:
	go build -o $(BINARY) goup.go

install: build
	sudo mv $(BINARY) $(INSTALL_DIR)

clean:
	rm -f $(BINARY)

.PHONY: build install clean
