SHELL :=/bin/bash
CWD := $(PWD)

build-dev:
	@docker build \
	-t app/auth \
    --target dev .
build-prod:
	@docker build \
	-t app/auth \
	--target prod .
run:
	@docker run \
	-p 4000:4000 \
  	--env-file=.env \
  	app/auth
	  
build-app:
	@exec go build

.PHONY: all
