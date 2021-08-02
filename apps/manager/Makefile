SHELL :=/bin/bash
CWD := $(PWD)

build-dev:
	@docker build \
	-t app/manager \
    --target dev .

build:
	@docker build \
	-t app/manager \
	--target prod .

run:
	@docker run \
	-p 4000:4000 \
  	--env-file=.env \
  	app/manager
	  
build-app:
	@exec go build

.PHONY: all
