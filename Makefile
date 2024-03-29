ifeq ($(OS),Windows_NT)
  PLATFORM="Windows"
else
  ifeq ($(shell uname),Darwin)
    PLATFORM="MacOS"
  else
    PLATFORM="Linux"
  endif
endif

all: build 
export GO111MODULE=on
export OUTPUT=./output/

build:
	PLATFORM=$(PLATFORM) ./build.sh

test:
	go test -cover ./...

clean:
	rm -rf $(OUTPUT)
	rm -f address
	rm -f private.key
	rm -f public.key
	rm -f mnemonic

.PHONY: all test clean
