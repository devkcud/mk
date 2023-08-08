all: build install

build:
	go build -o mk main.go

install:
	mv mk ${HOME}/.local/bin/mk
