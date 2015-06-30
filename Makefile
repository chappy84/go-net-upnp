###################################################################
#
# net/upnp
#
# Copyright (C) Satoshi Konno 2015
#
# This is licensed under BSD-style license, see file COPYING.
#
###################################################################

PREFIX?=$(shell pwd)
GOPATH=$(shell pwd)

packages = net/upnp net/upnp/log net/upnp/ssdp
	
all: build

format:
	gofmt -w src cmd

package: format $(shell find . -type f -name '*.go')
	go build -v ${packages}

${PREFIX}/bin/upnpdump: $(shell find ./cmd/upnpdump -type f -name '*.go')
	go build -o $@ ./cmd/upnpdump

build: ${PREFIX}/bin/upnpdump

test: package
	go test -v ${packages}

install: build
	go install ${packages}

clean:
	rm -rf _obj
	go clean -i ${packages}
